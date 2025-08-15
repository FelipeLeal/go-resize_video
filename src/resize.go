package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// Define constants for file directories and port.
const (
	uploadDir = "./uploads"
	outputDir = "./output"
	port      = ":8080"
)

// init runs before main() to ensure directories exist.
func init() {
	// os.ModePerm is a permission flag that grants read, write, and execute permissions.
	// os.MkdirAll creates the directories and any necessary parents.
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create upload directory: %v", err)
	}
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}
}

func main() {
	// Use gin.Default() for a router with default middleware (logger and recovery).
	router := gin.Default()

	// Use a custom CORS middleware to handle cross-origin requests.
	router.Use(corsMiddleware())

	// Define a group of API routes to organize the endpoints.
	api := router.Group("/api")
	{
		api.POST("/upload", uploadHandler)
		api.GET("/download/:fileName", downloadHandler)
	}

	// This is where a static file server would be, if you had a frontend in this repository.
	// For now, it's commented out as you'll have a separate frontend.
	// router.Static("/", "./static")

	fmt.Printf("Server running on http://localhost%s\n", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

// uploadHandler handles the POST request to upload and resize a video.
func uploadHandler(c *gin.Context) {
	// A struct to represent a successful JSON response.
	type UploadResponse struct {
		Message     string `json:"message"`
		DownloadURL string `json:"download_url"`
	}

	// Handle multipart form data. The video file should be named "video".
	file, err := c.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read video file from form data"})
		return
	}

	// Get the resolution from the form data.
	resolution := c.PostForm("resolution")
	if resolution == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Resolution is required"})
		return
	}

	// Create a new file in the upload directory.
	uploadPath := filepath.Join(uploadDir, file.Filename)
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open uploaded file"})
		return
	}
	defer src.Close()

	// Save the file to disk.
	dst, err := os.Create(uploadPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save uploaded file"})
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write file to disk"})
		return
	}

	// Define the output path for the resized video.
	outputPath := filepath.Join(outputDir, "resized_"+file.Filename)

	// Execute the FFmpeg command to resize the video.
	// The exec.Command function runs a command and its arguments.
	cmd := exec.Command("ffmpeg", "-i", uploadPath, "-vf", "scale="+resolution, outputPath)
	if err := cmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to resize video with FFmpeg"})
		return
	}

	// If successful, respond with a JSON object.
	downloadURL := fmt.Sprintf("/api/download/%s", "resized_"+file.Filename)
	c.JSON(http.StatusOK, UploadResponse{
		Message:     "Video resized successfully!",
		DownloadURL: downloadURL,
	})
}

// downloadHandler serves the resized video for download.
func downloadHandler(c *gin.Context) {
	// Get the file name from the URL path parameter.
	fileName := c.Param("fileName")
	if fileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File name is required"})
		return
	}

	// Construct the full path to the file.
	filePath := filepath.Join(outputDir, fileName)

	// Check if the file exists.
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// Serve the file as an attachment.
	c.FileAttachment(filePath, fileName)
}

// corsMiddleware is a custom Gin middleware to enable CORS.
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Allow requests from any origin. For production, you should restrict this.
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")

		// If the request is an OPTIONS preflight request, respond with 204 No Content.
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// Continue down the middleware chain.
		c.Next()
	}
}
