// # Ex49: Flickr Photo Search
//
// - Take in a search string via GUI.
// - Fetch Flickr’s public photo feed matching the search.
// - Display resulting photos visually.

package main

import (
	"encoding/json"
	"fmt"
	"image"
	"io"
	"net/http"
	"strings"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/nfnt/resize" // Add this import for resizing images
)

type FlickrResponse struct {
	Items []struct {
		Media struct {
			M string `json:"m"`
		} `json:"media"`
	} `json:"items"`
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Flickr Photo Search")

	tagEntry := widget.NewEntry()
	tagEntry.SetPlaceHolder("Enter tags...")
	tagEntryContainer := container.New(layout.NewGridWrapLayout(fyne.NewSize(400, tagEntry.MinSize().Height)), tagEntry)
	statusLabel := widget.NewLabel("Ready")
	imageContainer := container.NewGridWrap(fyne.NewSize(200, 200)) // Fixed grid size per image
	scrollableContainer := container.NewScroll(imageContainer)
	scrollableContainer.SetMinSize(fyne.NewSize(800, 400)) // Fixed scroll area size

	searchButton := widget.NewButton("Search", func() {
		tags := tagEntry.Text
		if tags != "" {
			updateImages(imageContainer, tags, statusLabel)
		}
	})
	toolbar := container.NewHBox(
		layout.NewSpacer(),
		widget.NewLabel("Tags:"),
		tagEntryContainer,
		searchButton,
	)

	statusBar := container.NewHBox(statusLabel)

	content := container.NewBorder(toolbar, statusBar, nil, nil, scrollableContainer)
	myWindow.SetContent(content)

	go updateImages(imageContainer, "", statusLabel)

	myWindow.Resize(fyne.NewSize(800, 600))
	fmt.Println("Starting the application...")
	myWindow.ShowAndRun()
}

func updateImages(imageContainer *fyne.Container, tags string, statusLabel *widget.Label) {
	statusLabel.SetText("Fetching")

	go func() {
		url := fmt.Sprintf("https://www.flickr.com/services/feeds/photos_public.gne?format=json&tags=%s", tags)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error: Failed to fetch data from Flickr")
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error: Failed to read Flickr response")
			return
		}

		var flickrData FlickrResponse
		jsonStr := strings.TrimPrefix(string(body), "jsonFlickrFeed(")
		jsonStr = strings.TrimSuffix(jsonStr, ")")
		if err := json.Unmarshal([]byte(jsonStr), &flickrData); err != nil {
			fmt.Println("Error: Failed to parse Flickr response")
			return
		}

		var wg sync.WaitGroup
		imageChan := make(chan fyne.CanvasObject, len(flickrData.Items))

		for _, item := range flickrData.Items {
			wg.Add(1)
			go func(mediaURL string) {
				defer wg.Done()

				imgResp, err := http.Get(mediaURL)
				if err != nil {
					fmt.Println("Error: Failed to fetch image:", err)
					return
				}
				defer imgResp.Body.Close()

				img, _, err := image.Decode(imgResp.Body)
				if err != nil {
					fmt.Println("Error: Failed to decode image:", err)
					return
				}

				resizedImg := resize.Thumbnail(200, 200, img, resize.Lanczos3)

				fyneImg := canvas.NewImageFromImage(resizedImg)
				fyneImg.SetMinSize(fyne.NewSize(200, 200))
				fyneImg.FillMode = canvas.ImageFillContain

				imageChan <- fyneImg
			}(item.Media.M)
		}

		// Wait for all goroutines to finish
		wg.Wait()
		close(imageChan)

		// Collect images from the channel
		var images []fyne.CanvasObject
		for img := range imageChan {
			images = append(images, img)
		}

		// Update the GUI
		imageContainer.Objects = images
		imageContainer.Refresh()

		if len(images) > 0 {
			statusLabel.SetText(fmt.Sprintf("Tags: %s", tags))
		} else {
			statusLabel.SetText(fmt.Sprintf("No images found for: %s", tags))
		}
	}()
}
