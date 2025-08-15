package main

import (
	"fmt"
	"log"

	"resize_video/src/config"
	"resize_video/src/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize configuration and create necessary directories.
	appConfig := config.NewConfig()
	if err := appConfig.Init(); err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// Create a new Gin router.
	router := gin.Default()

	// Use the CORS middleware from our handlers package.
	router.Use(handlers.CORS())

	// Define a group for our API endpoints.
	api := router.Group("/api")
	{
		api.POST("/upload", handlers.UploadHandler)
		api.GET("/download/:fileName", handlers.DownloadHandler)
	}

	// Start the server.
	fmt.Printf("Server running on http://localhost%s\n", appConfig.Port)
	if err := router.Run(appConfig.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
