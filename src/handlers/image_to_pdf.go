package handlers

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"resize_video/src/models"
)

// ImageToPDFHandler handles image uploads and creates a PDF with the images.
// @Summary      Convert images to PDF
// @Description  Uploads one or more image files and converts them into a single PDF document. The PDF can be downloaded via the returned URL.
// @Tags         pdf
// @Accept       multipart/form-data
// @Produce      json
// @Param        images      formData  file    true  "Image files to include in the PDF. Provide multiple 'images' parts for multiple files."
// @Success      200         {object}  models.PDFResponse
// @Failure      400         {object}  models.ErrorResponse
// @Failure      500         {object}  models.ErrorResponse
// @Router       /image-to-pdf [post]
func (h *Handlers) ImageToPDFHandler(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(400, models.ErrorResponse{Error: "Invalid form data"})
		return
	}
	files := form.File["images"]
	if len(files) == 0 {
		c.JSON(400, models.ErrorResponse{Error: "No images uploaded"})
		return
	}

	var imagePaths []string
	for _, file := range files {
		filename := filepath.Base(file.Filename)
		path := filepath.Join(h.AppConfig.UploadDir, filename)
		if err := c.SaveUploadedFile(file, path); err != nil {
			c.JSON(500, models.ErrorResponse{Error: fmt.Sprintf("Failed to save image: %s", filename)})
			return
		}
		imagePaths = append(imagePaths, path)
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	for _, imgPath := range imagePaths {
		pdf.AddPage()
		pdf.ImageOptions(imgPath, 10, 10, 190, 0, false, gofpdf.ImageOptions{ImageType: "", ReadDpi: true}, 0, "")
	}

	// Generate a unique name for the PDF to avoid overwrites.
	pdfName := fmt.Sprintf("images_%d.pdf", time.Now().UnixNano())
	pdfPath := filepath.Join(h.AppConfig.OutputDir, pdfName)
	if err := pdf.OutputFileAndClose(pdfPath); err != nil {
		c.JSON(500, models.ErrorResponse{Error: "Failed to create PDF"})
		return
	}

	// Return a download URL, not a file path.
	downloadURL := fmt.Sprintf("/api/download/%s", pdfName)
	c.JSON(200, models.PDFResponse{
		Message:     "PDF created successfully!",
		DownloadURL: downloadURL,
	})
}
