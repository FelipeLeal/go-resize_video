package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"resize_video/src/models"

	"github.com/gin-gonic/gin"
)

// DownloadHandler serves the resized video for download.
func (h *Handlers) DownloadHandler(c *gin.Context) {
	fileName := c.Param("fileName")
	if fileName == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "File name is required"})
		return
	}

	filePath := filepath.Join(h.AppConfig.OutputDir, fileName)

	// Check if the file exists.
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "File not found"})
		return
	}

	// Serve the file as an attachment.
	c.FileAttachment(filePath, fileName)
}
