package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"resize_video/src/models"

	"github.com/gin-gonic/gin"
)

// DownloadHandler serves the resized video for download.
// @Summary      Download a resized video
// @Description  Downloads a resized video file by its name.
// @Tags         videos
// @Produce      application/octet-stream
// @Param        fileName  path      string  true  "File name of the resized video"
// @Success      200       {file}    file    "The video file"
// @Failure      400       {object}  models.ErrorResponse
// @Failure      404       {object}  models.ErrorResponse
// @Router       /download/{fileName} [get]
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
