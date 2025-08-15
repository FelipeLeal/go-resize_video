package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	uploadDir   = "./uploads"
	outputDir   = "./output"
	port        = ":8080"
)

func main() {
	// Create directories for uploads and output
	os.MkdirAll(uploadDir, os.ModePerm)
	os.MkdirAll(outputDir, os.ModePerm)

	// Serve static files (HTML, CSS, JS)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Handle video upload and resizing
	http.HandleFunc("/upload", uploadHandler)

	// Serve resized video for download
	http.HandleFunc("/download", downloadHandler)

	fmt.Printf("Server running on http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the uploaded file
	file, header, err := r.FormFile("video")
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save the uploaded file
	uploadPath := filepath.Join(uploadDir, header.Filename)
	outFile, err := os.Create(uploadPath)
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}

	// Get the selected resolution
	resolution := r.FormValue("resolution")

	// Resize the video using FFmpeg
	outputPath := filepath.Join(outputDir, "resized_"+header.Filename)
	cmd := exec.Command("ffmpeg", "-i", uploadPath, "-vf", "scale="+resolution, outputPath)
	err = cmd.Run()
	if err != nil {
		http.Error(w, "Failed to resize video", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Video resized successfully! <a href='/download?file=%s'>Download</a>", "resized_"+header.Filename)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("file")
	if fileName == "" {
		http.Error(w, "File not specified", http.StatusBadRequest)
		return
	}

	filePath := filepath.Join(outputDir, fileName)
	http.ServeFile(w, r, filePath)
}
