package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// File repersents the structure of a file in the application.
type File struct {
	ID        string
	Name      string
	Path      string
	Thumbnail string
}

// files holds list of uploaded file.
var files = []File{}

func main() {
	r := gin.Default()
	//Serve static files from "static" directory.
	r.Static("/static", "./static")
	// Load HTML templates from "templates" directory.
	r.LoadHTMLGlob("templates/*.html")

	//route to handle
	r.GET("/", homeHandler)               //home page
	r.GET("/upload", uploadHandler)       //upload page
	r.GET("/:id", fileHandler)            //individual file display
	r.POST("/upload", uploadFilesHandler) //file uploads

	// Start the server on port 8080.
	r.Run(":8080")
}

func homeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"File": files,
	})
}

func uploadHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}

func fileHandler(c *gin.Context) {
	fileID := c.Param("id")

	//Find the file with matching id.
	var selectedFile File
	for _, file := range files {
		if file.ID == fileID {
			selectedFile = file
			break
		}
	}

	if selectedFile.ID == "" {
		c.String(http.StatusNotFound, "file not found.")
		return
	}

	c.HTML(http.StatusOK, "file.html", gin.H{
		"file": selectedFile,
	})
}

func uploadFilesHandler(c *gin.Context) {
	c.Request.ParseMultipartForm(10 << 20) //10MB max memory

	//Retrive file from form.
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	//Create the file path for saving the uploaded file
	filePath := filepath.Join("static", "uploads", header.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		c.String(http.StatusInternalServerError, "error saving the file")
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
