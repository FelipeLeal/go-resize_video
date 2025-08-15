package handlers

import (
	"net/http"

	"resize_video/src/config"
	"resize_video/src/service"

	"github.com/gin-gonic/gin"
)

// Handlers struct holds dependencies like config and services.
type Handlers struct {
	AppConfig    *config.AppConfig
	VideoService *service.VideoService
}

// NewHandlers creates a new Handlers instance with its dependencies.
func NewHandlers(cfg *config.AppConfig) *Handlers {
	return &Handlers{
		AppConfig:    cfg,
		VideoService: service.NewVideoService(cfg.UploadDir, cfg.OutputDir),
	}
}

// CORS is a custom Gin middleware to enable CORS.
func (h *Handlers) CORS() gin.HandlerFunc {
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