package models

// UploadResponse represents the JSON response for a successful upload.
type UploadResponse struct {
	Message     string `json:"message"`
	DownloadURL string `json:"download_url"`
}

// ErrorResponse represents the JSON response for an error.
type ErrorResponse struct {
	Error string `json:"error"`
}
