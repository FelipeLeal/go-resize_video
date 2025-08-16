# 🎥 Go Video Resizer API

A modern, Docker-based API for resizing videos, built with Go, Gin, and FFmpeg. Features a hot-reload development environment powered by Air.

## ✨ Features

- **RESTful API**: Simple endpoints for uploading and downloading videos.
- **Powerful Video Processing**: Leverages FFmpeg for resizing and format handling.
- **Multiple Formats**: Supports MP4, AVI, MOV, MKV, and more video formats.
- **Configurable Resolutions**: Easily define target resolutions for resizing.
- **Hot Reload**: Development environment with `air` for automatic recompilation on code changes.
- **Dockerized**: Fully containerized with Go and FFmpeg pre-installed for a consistent environment.

## 🚀 Quick Start

### Prerequisites

- Docker and Docker Compose installed on your system.
- Git (to clone the repository).

### Installation & Running

1. **Clone the repository**
   ```bash
   git clone <your-repository-url>
   cd <your-repository-directory>
   ```

2. **Start the application**
   ```bash
   docker-compose up --build
   ```

3. **Access the API**

   The server will be running and accessible. You can now send requests to the API endpoints, for example, using `curl` or Postman. The server listens on `http://localhost:8080`.

### Example API Usage

**Upload and resize a video:**

```bash
curl -X POST http://localhost:8080/api/upload \
  -F "file=@/path/to/your/video.mp4" \
  -F "resolution=1280x720"
```

**Download the resized video (using the filename from the upload response):**

```bash
curl -o resized_video.mp4 http://localhost:8080/api/download/your-resized-filename.mp4
```

## 🛠️ Development

The project uses Air for live reloading. When you run `docker-compose up`, Air monitors your `.go` files for changes. Upon detecting a change, it automatically rebuilds and restarts the application inside the container.

You can configure Air's behavior in the `.air.toml` file.

## 🔧 Technical Details

### Architecture

- **Backend**: Go with the Gin web framework.
- **Video Processing**: FFmpeg command-line tool.
- **Container**: `golang:1.24-alpine` base image with FFmpeg installed.
- **Live Reload**: Air for hot-reloading during development.

### API Endpoints

- `POST /api/upload`: Handles video upload and queues it for resizing. Requires a multipart form with `file` and `resolution` fields.
- `GET /api/download/:fileName`: Serves the resized video file for download.

### Project Structure

```
.
├── .air.toml               # Configuration for Air (hot-reloading)
├── .devcontainer/
│   └── devcontainer.json   # VS Code Dev Container configuration
├── Dockerfile              # Defines the application's Docker image
├── docker-compose.yml      # Defines the services, networks, and volumes
├── go.mod                  # Go module definitions
├── go.sum                  # Go module checksums
├── src/
│   ├── config/             # Application configuration
│   ├── handlers/           # Gin HTTP handlers (controllers)
│   ├── main.go             # Application entry point
│   ├── models/             # Data structures (e.g., API responses)
│   └── services/           # Business logic (e.g., video processing)
└── README.md               # This file
```

## 🐳 Docker Configuration

### Base Image

- `golang:1.24-alpine`: A lightweight official Go image based on Alpine Linux.

### Runtime Dependencies

- **Go**: For building and running the web server
- **FFmpeg**: For all video processing tasks.
- **Air**: For live-reloading during development.

### Volumes

- `.:/app`: Mounts the entire project directory into the container's working directory, enabling live code changes.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 📄 License

This project is licensed under the terms specified in the LICENSE file.

## 🔧 Troubleshooting

**Container won't start:**
- Ensure Docker and Docker Compose are running.
- Check if port 8080 is available on your host machine.
- Run `docker-compose down` followed by `docker-compose up --build` to ensure a clean start.

**Video processing fails:**
- Verify the input file is a valid and supported video format.
- Check the container logs for FFmpeg errors using `docker-compose logs`.

**Hot reload not working:**
- Ensure you are running the service with `docker-compose up`.
- Verify that the volume `.:/app` is correctly defined in `docker-compose.yml`.
- Check the container logs; Air will print messages when it detects file changes and restarts the server.
