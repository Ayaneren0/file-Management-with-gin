package main

import "github.com/gin-gonic/gin"

//File repersents the structure of a file in the application.
type File struct {
	ID        string
	Name      string
	Path      string
	Thumbnail string
}

//files holds list of uploaded file.
var files = &File{}

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

}

func uploadHandler(c *gin.Context) {

}

func fileHandler(c *gin.Context) {

}

func uploadFilesHandler(c *gin.Context) {

}
