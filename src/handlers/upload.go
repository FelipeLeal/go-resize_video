package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"resize_video/src/models"

	"github.com/gin-gonic/gin"
)

// UploadHandler handles the POST request to upload and resize a video.
// @Summary      Upload and resize a video
// @Description  Uploads a video file and a desired resolution, resizes it, and returns a download link.
// @Tags         videos
// @Accept       multipart/form-data
// @Produce      json
// @Param        video       formData  file    true  "Video file to upload"
// @Param        resolution  formData  string  true  "Desired resolution (e.g., 640x480)"
// @Success      200         {object}  models.UploadResponse
// @Failure      400         {object}  models.ErrorResponse
// @Failure      500         {object}  models.ErrorResponse
// @Router       /upload [post]
func (h *Handlers) UploadHandler(c *gin.Context) {
	file, err := c.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Unable to read video file from form data"})
		return
	}

	resolution := c.PostForm("resolution")
	if resolution == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Resolution is required"})
		return
	}

	// Save the file to the upload directory.
	uploadPath := filepath.Join(h.AppConfig.UploadDir, file.Filename)
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to open uploaded file"})
		return
	}
	defer src.Close()

	dst, err := os.Create(uploadPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to save uploaded file"})
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to write file to disk"})
		return
	}

	// Call the video service to resize the video.
	outputPath, err := h.VideoService.ResizeVideo(uploadPath, resolution)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	// Return a JSON response with the download URL.
	downloadURL := fmt.Sprintf("/api/download/%s", filepath.Base(outputPath))
	c.JSON(http.StatusOK, models.UploadResponse{
		Message:     "Video resized successfully!",
		DownloadURL: downloadURL,
	})
}
