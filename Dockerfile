# Use the same base image with FFmpeg and Alpine Linux
FROM jrottenberg/ffmpeg:6.1-alpine

# Set the working directory inside the container
WORKDIR /app

# Install Go and entr, which is used by the dev.sh script for hot-reloading
RUN apk update && \
    apk add --no-cache go entr

# Copy the application source code.
COPY src/resize.go .

# Initialize the Go module and get dependencies.
# The `go mod init` and `go get` commands will create go.mod and go.sum.
# This ensures that the dev.sh script can find them.
RUN go mod init resize_video && \
    go get github.com/gin-gonic/gin && \
    go mod tidy

# Copy the development script
COPY src/dev.sh /usr/local/bin/dev.sh

# Fix line endings and make the script executable
RUN sed -i 's/\r$//' /usr/local/bin/dev.sh && chmod +x /usr/local/bin/dev.sh

# Expose the port the app runs on
EXPOSE 8080

# The entrypoint is the dev.sh script, which handles the hot-reloading of the Go server
CMD ["/usr/local/bin/dev.sh"]