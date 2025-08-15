#!/bin/sh

# Development script with hot reload functionality for the refactored Go API
# This script is designed to run from the project's root directory.

echo "Application started with PID $$"

# Function to calculate a hash of all Go files and module files
calculate_hash() {
    # Find all .go files recursively, and include the go.mod and go.sum files.
    find . -name '*.go' -o -name 'go.mod' -o -name 'go.sum' -exec cat {} \; | md5sum | cut -d' ' -f1
}

# Function to start the application
start_app() {
    # Check for the compiled binary and run it
    if [ -f "./server" ]; then
        ./server &
        APP_PID=$!
        echo "Server started at :8080"
    else
        echo "Server executable not found. Please check the build process."
    fi
}

# Function to build the application
build_app() {
    echo ">>> Changes detected! Building and restarting application..."
    
    # Ensure go.mod exists and dependencies are up-to-date
    if [ ! -f "go.mod" ]; then
        echo "go.mod not found, please run 'go mod init' and 'go get' manually."
        return 1
    fi
    go mod tidy

    # Build the main package in the current directory and create a binary named 'server'
    go build -o server .

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

# Cleanup function on exit
cleanup() {
    echo "Cleaning up..."
    stop_app
    exit 0
}

# Set up signal handlers for graceful shutdown
trap cleanup SIGINT SIGTERM

echo "Watching for file changes..."

# Calculate initial hash
current_hash=$(calculate_hash)
echo "Initial hash: $current_hash"

# Perform initial build and start
if build_app; then
    start_app
fi

# Watch for changes in a continuous loop
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