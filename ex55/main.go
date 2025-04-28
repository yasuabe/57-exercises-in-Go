// # Ex55: Text Sharing
//
// - Create a web app for sharing text snippets (like Pastie).
// - Users can enter and save text through a form.
// - The app stores the text in a persistent data store.
// - Each saved text is assigned a URL-safe slug (e.g. via hash like SHA or MD5), not a primary key.
// - Users can:
//   - View the text by visiting its unique URL.
//   - Click "Edit" to copy it into the text submission form again.

package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/net/context"
)

type Snippet struct {
	Slug    string `bson:"slug"`
	Content string `bson:"content"`
}

var (
	client     *mongo.Client
	collection *mongo.Collection
	templates  *template.Template
)

func main() {
	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	// Ping MongoDB to verify connection
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	// Set up the collection
	collection = client.Database("text_sharing").Collection("snippets")

	// Parse templates
	templates = template.Must(template.ParseFiles("ex55/display.html", "ex55/input.html"))

	// Set up routes
	http.HandleFunc("/ex55", handleForm)
	http.HandleFunc("/ex55/", handleView)
	http.HandleFunc("/ex55/edit", handleEdit)

	// Start the server
	fmt.Println("Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Render the input form
		templates.ExecuteTemplate(w, "input.html", nil)
		return
	}

	if r.Method == http.MethodPost {
		// Process form submission
		snippet := r.FormValue("snippet")
		if snippet == "" {
			http.Error(w, "Snippet cannot be empty", http.StatusBadRequest)
			return
		}

		// Generate slug
		slug := generateSlug(snippet)

		// Save to MongoDB
		_, err := collection.InsertOne(context.TODO(), Snippet{
			Slug:    slug,
			Content: snippet,
		})
		if err != nil {
			http.Error(w, "Failed to save snippet", http.StatusInternalServerError)
			return
		}

		// Redirect to the view page
		http.Redirect(w, r, fmt.Sprintf("/ex55/%s", slug), http.StatusSeeOther)
	}
}

func handleView(w http.ResponseWriter, r *http.Request) {
	// Extract slug from URL
	slug := r.URL.Path[len("/ex55/"):]

	// Fetch snippet from MongoDB
	var snippet Snippet
	err := collection.FindOne(context.TODO(), bson.M{"slug": slug}).Decode(&snippet)
	if err != nil {
		http.Error(w, "Snippet not found", http.StatusNotFound)
		return
	}

	// Render the display template with the snippet data
	err = templates.ExecuteTemplate(w, "display.html", snippet)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

func handleEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Extract slug from form
	slug := r.FormValue("slug")

	// Fetch snippet from MongoDB
	var snippet Snippet
	err := collection.FindOne(context.TODO(), bson.M{"slug": slug}).Decode(&snippet)
	if err != nil {
		http.Error(w, "Snippet not found", http.StatusNotFound)
		return
	}

	// Render the input form with the existing snippet
	templates.ExecuteTemplate(w, "input.html", snippet)
}

func generateSlug(snippet string) string {
	// Generate a unique slug using MD5 hash
	text := fmt.Sprintf("%s%s", snippet, uuid.New().String())
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
