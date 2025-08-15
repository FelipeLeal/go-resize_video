package service

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

// VideoService handles all video-related operations.
type VideoService struct {
	UploadDir string
	OutputDir string
}

// NewVideoService creates a new instance of the VideoService.
func NewVideoService(uploadDir, outputDir string) *VideoService {
	return &VideoService{
		UploadDir: uploadDir,
		OutputDir: outputDir,
	}
}

// ResizeVideo uses FFmpeg to resize a video file.
func (s *VideoService) ResizeVideo(inputPath, resolution string) (string, error) {
	// Create the output path for the resized video.
	fileName := filepath.Base(inputPath)
	outputPath := filepath.Join(s.OutputDir, fmt.Sprintf("resized_%s", fileName))

	// Execute the FFmpeg command.
	cmd := exec.Command("ffmpeg", "-i", inputPath, "-vf", "scale="+resolution, outputPath)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to resize video: %w", err)
	}

	return outputPath, nil
}