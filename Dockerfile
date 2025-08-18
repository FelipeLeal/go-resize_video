FROM golang:1.24-alpine

# Install necessary packages:
# - ffmpeg for video processing
# - bash for a better interactive shell
# - git for go modules (required by `go install`)
# - ca-certificates for https requests
RUN apk update && apk add --no-cache \
    bash \
    ca-certificates \
    ffmpeg \
    git

# Install Air for hot-reloading. The binary will be in the PATH.
# Pinning the version of Air ensures reproducible builds.
RUN go install github.com/air-verse/air@v1.62.0

# Install Swag for API documentation generation.
# Pinning the version ensures reproducible builds.
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.3

# Set the working directory inside the container.
# With the volume mount in docker-compose, this will be the project root.
WORKDIR /app

# The command will be provided by docker-compose.yml (i.e., ["air"])
# We expose the port here for documentation.
EXPOSE 8080
