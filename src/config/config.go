package config

import (
	"log"
	"os"
)

// AppConfig holds the application's configuration settings.
type AppConfig struct {
	UploadDir string
	OutputDir string
	Port      string
}

// NewConfig creates and returns a new AppConfig instance.
func NewConfig() *AppConfig {
	return &AppConfig{
		UploadDir: "./uploads",
		OutputDir: "./output",
		Port:      ":8080",
	}
}

// Init creates the necessary directories if they don't exist.
func (c *AppConfig) Init() error {
	log.Println("Initializing application directories...")
	if err := os.MkdirAll(c.UploadDir, os.ModePerm); err != nil {
		return err
	}
	if err := os.MkdirAll(c.OutputDir, os.ModePerm); err != nil {
		return err
	}
	return nil
}
