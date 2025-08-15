#!/bin/sh

# Development script with hot reload functionality for Gin-based API
# Uses hash-based file watching to avoid infinite restart loops

echo "Application started with PID $$"

# Function to calculate hash of all Go files and go.mod
calculate_hash() {
    find . -name '*.go' -o -name 'go.mod' -o -name 'go.sum' -exec cat {} \; | md5sum | cut -d' ' -f1
}

# Function to start the application
start_app() {
    # Check if the 'server' executable exists and run it.
    if [ -f "./server" ]; then
        ./server &
        APP_PID=$!
        echo "Server started at :8080"
    else
        echo "Server executable not found. Please check build process."
    fi
}

# Function to build the application
build_app() {
    echo ">>> Changes detected! Building and restarting application..."
    # Ensure dependencies are up-to-date
    go mod tidy

    # Build the Go application into a binary named 'server'
    # Use -o to specify the output filename
    go build -o server resize.go

    if [ $? -ne 0 ]; then
        echo "Build failed. Server will not be restarted."
        return 1
    fi
    return 0
}

# Function to stop the application
stop_app() {
    if [ ! -z "$APP_PID" ]; then
        echo "Stopping application with PID $APP_PID"
        kill $APP_PID 2>/dev/null
        wait $APP_PID 2>/dev/null
    fi
}

# Cleanup function
cleanup() {
    echo "Cleaning up..."
    stop_app
    exit 0
}

# Set up signal handlers
trap cleanup SIGINT SIGTERM

echo "Watching for file changes..."

# Calculate initial hash
current_hash=$(calculate_hash)
echo "Initial hash: $current_hash"

# Perform initial build and start
if build_app; then
    start_app
fi

# Watch for changes
while true
do
    sleep 2
    new_hash=$(calculate_hash)
    
    if [ "$new_hash" != "$current_hash" ]; then
        stop_app
        current_hash=$new_hash
        if build_app; then
            start_app
        fi
    fi
done