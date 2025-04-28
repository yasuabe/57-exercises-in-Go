// # Ex54: URL Shortener
//
// - Create a web app that shortens long URLs (like goo.gl).
// - Features:
//   - A form to submit a long URL.
//   - Generates and stores a short URL (e.g. /abc1234) that redirects to the long one.
//   - Tracks how many times the short URL is visited.
//   - Provides a stats page (/abc1234/stats) showing:
//     - The short URL
//     - The original long URL
//     - Visit count
// Constraints:
// - Must use a persistent, shareable data store (e.g. DB, not memory).
// - Must validate that the input is a valid URL.

package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"regexp"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var (
	rdb    *redis.Client
	ctx    = context.Background()
	tmpl   = template.Must(template.ParseGlob("ex54/*.html"))
	urlReg = regexp.MustCompile(`^(http|https)://[^\s/$.?#].[^\s]*$`)
)

func main() {
	// Initialize Redis client
	rdb = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	// Initialize router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/ex54/", handleForm).Methods("GET")
	r.HandleFunc("/ex54/", handleShorten).Methods("POST")
	r.HandleFunc("/ex54/{shortURL:[a-zA-Z0-9]+}", handleRedirect).Methods("GET")
	r.HandleFunc("/ex54/{shortURL:[a-zA-Z0-9]+}/stats", handleStats).Methods("GET")

	// Start the server
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	tmpl.ExecuteTemplate(w, "input.html", nil)
}

func handleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	longURL := r.FormValue("long_url")
	if !urlReg.MatchString(longURL) {
		tmpl.ExecuteTemplate(w, "input.html", map[string]string{"Error": "Invalid URL"})
		return
	}

	// Increment sequence number in Redis
	seq, err := rdb.Incr(ctx, "ex54:short_urls:sequence").Result()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Generate short URL
	shortURL := makeShortURL(seq)

	// Store the mapping in Redis
	err = rdb.HSet(ctx, "ex54:short_urls:"+shortURL, map[string]interface{}{
		"long_url":    longURL,
		"visit_count": 0,
	}).Err()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Redirect to stats page
	http.Redirect(w, r, "/ex54/"+shortURL+"/stats", http.StatusSeeOther)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	// Get the long URL from Redis
	longURL, err := rdb.HGet(ctx, "ex54:short_urls:"+shortURL, "long_url").Result()
	if err == redis.Nil {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Increment the visit count
	rdb.HIncrBy(ctx, "ex54:short_urls:"+shortURL, "visit_count", 1)

	// Redirect to the long URL
	http.Redirect(w, r, longURL, http.StatusFound)
}

func handleStats(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	// Get data from Redis
	data, err := rdb.HGetAll(ctx, "ex54:short_urls:"+shortURL).Result()
	if err == redis.Nil {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Render the stats page
	tmpl.ExecuteTemplate(w, "stats.html", map[string]interface{}{
		"LongURL":    data["long_url"],
		"ShortURL":   shortURL,
		"VisitCount": data["visit_count"],
		"Origin":     r.Host,
	})
}

func makeShortURL(seq int64) string {
	const base62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var shortURL []byte
	num := seq + 10_000_000_000
	for num > 0 {
		remainder := num % 62
		shortURL = append(shortURL, base62[remainder])
		num /= 62
	}
	// Reverse the slice to get the correct order
	for i, j := 0, len(shortURL)-1; i < j; i, j = i+1, j-1 {
		shortURL[i], shortURL[j] = shortURL[j], shortURL[i]
	}
	return string(shortURL)
}
