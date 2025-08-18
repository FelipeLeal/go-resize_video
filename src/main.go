package main

import (
	"fmt"
	"log"
	"net/http"

	"resize_video/src/config"
	"resize_video/src/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "resize_video/docs"
)

// @title           Resize Video API
// @version         1.0
// @description     This is a service to resize videos.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.example.com/support
// @contact.email  f3lipel3al@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api
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

	// Add Swagger endpoint.
	// This adds a redirect for users who navigate to /swagger without a trailing slash.
	router.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	// The gin-swagger handler serves the Swagger UI. We wrap it to handle the redirect for the root path.
	swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)
	router.GET("/swagger/*any", func(c *gin.Context) {
		if c.Param("any") == "/" {
			c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
			return
		}
		swaggerHandler(c)
	})

	// Start the server.
	fmt.Printf("Server running on http://localhost%s\n", appConfig.Port)
	if err := router.Run(appConfig.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
