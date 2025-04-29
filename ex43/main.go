// # Ex43: Website Generator
//
// - Prompt for:
//   - Site name
//   - Author name
//   - Whether to create js/ and css/ folders
// - Generate:
//   - A site root folder named after the site
//   - index.html with <title> and <meta> using input
//   - Optionally: js/ and/or css/ subfolders
// - Output messages confirming each created file/folder.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for site name
	fmt.Print("Site name: ")
	siteName, _ := reader.ReadString('\n')
	siteName = strings.TrimSpace(siteName)

	// Prompt for author name
	fmt.Print("Author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	// Prompt for JavaScript folder
	fmt.Print("Do you want a folder for JavaScript? (y/n): ")
	jsFolder, _ := reader.ReadString('\n')
	jsFolder = strings.TrimSpace(jsFolder)

	// Prompt for CSS folder
	fmt.Print("Do you want a folder for CSS? (y/n): ")
	cssFolder, _ := reader.ReadString('\n')
	cssFolder = strings.TrimSpace(cssFolder)

	// Create the site root folder
	err := os.Mkdir(siteName, 0755)
	if err != nil {
		fmt.Printf("Error creating site folder: %v\n", err)
		return
	}
	fmt.Printf("Created ./%s\n", siteName)

	// Create index.html
	indexFilePath := fmt.Sprintf("./%s/index.html", siteName)
	indexFile, err := os.Create(indexFilePath)
	if err != nil {
		fmt.Printf("Error creating index.html: %v\n", err)
		return
	}
	defer indexFile.Close()

	// Write basic HTML content to index.html
	htmlContent := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="author" content="%s">
    <title>%s</title>
</head>
<body>
</body>
</html>`, author, siteName)
	indexFile.WriteString(htmlContent)
	fmt.Printf("Created %s\n", indexFilePath)

	// Optionally create js/ folder
	if strings.ToLower(jsFolder) == "y" {
		jsFolderPath := fmt.Sprintf("./%s/js", siteName)
		err := os.Mkdir(jsFolderPath, 0755)
		if err != nil {
			fmt.Printf("Error creating JavaScript folder: %v\n", err)
			return
		}
		fmt.Printf("Created %s/\n", jsFolderPath)
	}

	// Optionally create css/ folder
	if strings.ToLower(cssFolder) == "y" {
		cssFolderPath := fmt.Sprintf("./%s/css", siteName)
		err := os.Mkdir(cssFolderPath, 0755)
		if err != nil {
			fmt.Printf("Error creating CSS folder: %v\n", err)
			return
		}
		fmt.Printf("Created %s/\n", cssFolderPath)
	}
}
