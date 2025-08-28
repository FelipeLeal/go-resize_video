package models

// ErrorResponse defines the structure for a generic error response.
type ErrorResponse struct {
	Error string `json:"error"`
}

// UploadResponse defines the structure for a successful video upload response.
type UploadResponse struct {
	Message     string `json:"message"`
	DownloadURL string `json:"downloadURL"`
}

// PDFResponse defines the structure for a successful PDF creation response.
type PDFResponse struct {
	Message     string `json:"message"`
	DownloadURL string `json:"downloadURL"`
}
