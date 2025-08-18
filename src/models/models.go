package models

// UploadResponse defines the structure for a successful upload response.
type UploadResponse struct {
	Message     string `json:"message"`
	DownloadURL string `json:"downloadURL"`
}

// ErrorResponse defines the structure for a generic error response.
type ErrorResponse struct {
	Error string `json:"error"`
}
