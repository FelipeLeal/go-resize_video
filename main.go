package main

import (
	"fmt"
	"log"

	"resize_video/config"
	"resize_video/handlers"

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

	// Create a handler instance with the initialized config.
	h := handlers.NewHandlers(appConfig)

	// Use the CORS middleware from our handlers package.
	router.Use(h.CORS())

	// Define a group for our API endpoints.
	api := router.Group("/api")
	{
		api.POST("/upload", h.UploadHandler)
		api.GET("/download/:fileName", h.DownloadHandler)
	}

	// Start the server.
	fmt.Printf("Server running on http://localhost%s\n", appConfig.Port)
	if err := router.Run(appConfig.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
