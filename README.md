# 🎥 Video Resizer

A modern, Docker-based web application for resizing videos built with Go and FFmpeg. Features a sleek user interface with hot reload development environment.

## ✨ Features

- **Easy Video Resizing**: Upload videos and resize them to popular resolutions
- **Multiple Formats**: Supports MP4, AVI, MOV, MKV, and more video formats
- **Pre-defined Resolutions**: 
  - 📱 360p (640x360) - Mobile
  - 💻 720p HD (1280x720) - Standard
  - 🖥️ 1080p Full HD (1920x1080) - High Quality
  - 🎬 1440p QHD (2560x1440) - Premium
- **Modern UI**: Beautiful, responsive web interface with gradients and animations
- **Hot Reload**: Development environment with automatic restart on code changes
- **Docker Integration**: Fully containerized with FFmpeg pre-installed

## 🚀 Quick Start

### Prerequisites

- Docker and Docker Compose installed on your system
- Git (to clone the repository)

### Installation & Running

1. **Clone the repository**
   ```bash
   git clone https://github.com/FelipeLeal/go-resize_video.git
   cd go-resize_video
   ```

2. **Start the application**
   ```bash
   make up
   ```
   
   Or using docker-compose directly:
   ```bash
   docker-compose up
   ```

3. **Access the application**
   
   Open your browser and navigate to: `http://localhost:8080`

4. **Upload and resize videos**
   - Select a video file
   - Choose your desired resolution
   - Click "Upload and Resize Video"
   - Download the resized video

## 🛠️ Development

The project includes a hot reload development environment that automatically restarts the application when Go files are modified.

### Available Commands

```bash
# Start the application
make up

# Stop the application
make down

# View logs
make logs

# Restart the application
make restart

# Access the container shell
make shell

# View all containers
make ps

# Clean up Docker resources
make prune

# View all available commands
make help
```

### Project Structure

```
resize_video/
├── src/
│   ├── resize.go           # Main Go application
│   ├── dev.sh             # Hot reload development script
│   ├── input.tmp          # Temporary upload file
│   └── output.mp4         # Processed video output
├── Dockerfile             # Container configuration
├── docker-compose.yml     # Multi-container setup
├── Makefile              # Development commands
├── LICENSE               # Project license
└── README.md            # This file
```

## 🔧 Technical Details

### Architecture

- **Backend**: Go HTTP server with standard library
- **Video Processing**: FFmpeg for video conversion and resizing
- **Frontend**: HTML5 with embedded CSS and modern styling
- **Container**: Alpine Linux with FFmpeg pre-installed
- **Development**: Hash-based file watching for hot reload

### API Endpoints

- `GET /` - Upload form (main interface)
- `POST /upload` - Handle video upload and processing
- `GET /download` - Download processed video

### File Watching System

The development environment uses a sophisticated hash-based file watching system that:
- Monitors Go source files for changes
- Calculates MD5 hash of file contents
- Only restarts when actual content changes (not just timestamps)
- Provides zero false-positive restarts

## 🎨 UI Features

- **Responsive Design**: Works on desktop and mobile devices
- **Modern Styling**: CSS gradients, shadows, and smooth animations
- **File Upload**: Drag-and-drop style file input
- **Progress Feedback**: Clear status messages and success screens
- **Download Integration**: Direct download links for processed videos

## ⚙️ Configuration

### Supported Video Formats

The application leverages FFmpeg's extensive format support:
- **Input**: MP4, AVI, MOV, MKV, WebM, FLV, WMV, and more
- **Output**: MP4 (H.264 video, original audio codec)

### Resolution Options

Pre-configured resolution presets:
- **360p**: Optimized for mobile viewing and bandwidth conservation
- **720p HD**: Standard high-definition, good balance of quality and file size
- **1080p Full HD**: High quality for desktop viewing
- **1440p QHD**: Premium quality for high-resolution displays

## 🐳 Docker Configuration

### Base Image
- `jrottenberg/ffmpeg:6.1-alpine` - Provides FFmpeg 6.1 on Alpine Linux

### Runtime Dependencies
- **Go**: For building and running the web server
- **FFmpeg**: For video processing capabilities
- **entr**: For file watching and hot reload functionality

### Volumes
- `./src:/app` - Source code hot reload
- `./src/dev.sh:/usr/local/bin/dev.sh` - Development script

## 📝 Usage Examples

### Basic Video Resize
1. Upload a video file (any supported format)
2. Select "720p HD (1280x720)" from the dropdown
3. Click "Upload and Resize Video"
4. Wait for processing to complete
5. Click "Download Video" to get the resized file

### Development Workflow
1. Run `make up` to start the development environment
2. Edit Go files in the `src/` directory
3. Changes are automatically detected and the server restarts
4. Refresh your browser to see the changes

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 📄 License

This project is licensed under the terms specified in the LICENSE file.

## 🔧 Troubleshooting

### Common Issues

**Container won't start:**
- Ensure Docker is running
- Check if port 8080 is available
- Try `make down` followed by `make up`

**Video processing fails:**
- Verify the input file is a valid video format
- Check Docker logs with `make logs`
- Ensure sufficient disk space for temporary files

**Hot reload not working:**
- Check that volumes are properly mounted
- Verify file permissions on `dev.sh`
- Look for hash change messages in logs

### Performance Tips

- Use smaller input files for faster processing
- Choose appropriate resolution based on your needs
- Monitor Docker resource usage for large files

## 📞 Support

For issues, questions, or contributions, please:
1. Check existing issues in the repository
2. Create a new issue with detailed description
3. Include relevant logs and system information
