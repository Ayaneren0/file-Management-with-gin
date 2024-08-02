package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// File represents the structure of a file in the application.
type File struct {
	ID        string
	Name      string
	Path      string
	Thumbnail string
}

// files holds the list of uploaded files.
var files = []File{}

func main() {
	r := gin.Default()

	// Serve static files from the "static" directory.
	r.Static("/static", "./static")

	// Load HTML templates from the "templates" directory.
	r.LoadHTMLGlob("templates/*")

	// Route to serve the home page.
	r.GET("/", indexHandler)

	// Route to serve the file upload page.
	r.GET("/upload", uploadHandler)

	// Route to handle individual file display.
	r.GET("/files/:id", fileHandler)

	// Route to handle file uploads.
	r.POST("/upload", uploadFileHandler)

	// Start the server on port 8080.
	r.Run(":8080")
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Files": files,
	})
}
func uploadHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}

// fileHandler handles requests to view individual files by ID.
func fileHandler(c *gin.Context) {
	fileId := c.Param("id")

	// Find the file with the matching ID.
	var selectedFile File
	for _, file := range files {
		if file.ID == fileId {
			selectedFile = file
			break
		}
	}

	// If no file is found, return a 404 error.
	if selectedFile.ID == "" {
		c.String(http.StatusNotFound, "file ID not found!")
		return
	}

	// Render the file details page.
	c.HTML(http.StatusOK, "file.html", gin.H{
		"File": selectedFile,
	})
}

// uploadHandler handles file upload requests.
func uploadFileHandler(c *gin.Context) {
	c.Request.ParseMultipartForm(10 << 20) // 10 MB max memory

	// Retrieve the file from the form.
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Error retrieving the file")
		return
	}
	defer file.Close()

	// Create the file path for saving the uploaded file.
	filePath := filepath.Join("static", "uploads", header.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error saving the file")
		return
	}
	defer dst.Close()

	// Write the uploaded file to the destination.
	if _, err := dst.ReadFrom(file); err != nil {
		c.String(http.StatusInternalServerError, "Error writing the file")
		return
	}

	// Create a new File struct and add it to the files slice.
	newFile := File{
		ID:        header.Filename,
		Name:      header.Filename,
		Path:      filePath,
		Thumbnail: header.Filename, // Simplified; generate a real thumbnail in practice
	}
	files = append(files, newFile)

	// Redirect to the home page after successful upload.
	c.Redirect(http.StatusSeeOther, "/")
}
