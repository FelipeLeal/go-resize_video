package handlers

import (
	"net/http"

	"resize_video/src/config"
	"resize_video/src/service"

	"github.com/gin-gonic/gin"
)

// NewAppConfig creates a config instance for use in handlers.
var appConfig = config.NewConfig()
var videoService = service.NewVideoService(appConfig.UploadDir, appConfig.OutputDir)

// CORS is a custom Gin middleware to enable CORS.
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}