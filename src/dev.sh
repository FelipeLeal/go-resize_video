#!/bin/bash

# Development script with hot reload functionality
# Uses hash-based file watching to avoid infinite restart loops

echo "Application started with PID $$"

# Function to calculate hash of all Go files
calculate_hash() {
    find . -name '*.go' -exec cat {} \; | md5sum | cut -d' ' -f1
}

# Function to start the application
start_app() {
    go run resize.go &
    APP_PID=$!
    echo "Server started at :8080"
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

# Start the application
start_app

# Watch for changes
while true; do
    sleep 2
    new_hash=$(calculate_hash)
    
    if [ "$new_hash" != "$current_hash" ]; then
        echo ">>> Changes detected! Restarting application..."
        echo "Old hash: $current_hash"
        echo "New hash: $new_hash"
        
        stop_app
        current_hash=$new_hash
        start_app
    fi
done