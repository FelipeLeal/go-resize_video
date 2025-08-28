#!/bin/sh

# Exit immediately if a command exits with a non-zero status.
# Use 'set -eu' for POSIX-compliant error handling. 'pipefail' is a bash-specific option.
set -eu

# If no command is given, show usage.
if [ $# -eq 0 ]; then
    echo "Usage: dev <command>"
    echo ""
    echo "Available commands:"
    echo "  up             - Start the application with hot-reload (air)."
    echo "  tidy           - Sync Go modules (go mod tidy)."
    echo "  dl             - Download Go modules (go mod download)."
    echo "  swag           - Generate Swagger documentation (swag init)."
    echo "  init           - Install Go development tools (air, swag)."
    echo "  test           - Run Go tests (e.g., 'dev test ./...')."
    exit 1
fi

COMMAND=$1
shift

case "$COMMAND" in
    up)
        echo "--- Starting application with hot-reload (air) ---"
        # Ensure tools are installed before trying to run them
        command -v air >/dev/null 2>&1 || { echo >&2 "Error: 'air' is not installed. Please run 'dev install-tools' first."; exit 1; }
        command -v swag >/dev/null 2>&1 || { echo >&2 "Error: 'swag' is not installed. Please run 'dev install-tools' first."; exit 1; }
        # Generate swagger docs before starting
        # Use -d to specify the directory to parse for Go files.
        swag init -g main.go -d ./src
        exec air
        ;;
    tidy)
        echo "--- Syncing Go modules (go mod tidy) ---"
        exec go mod tidy
        ;;
    dl)
        echo "--- Downloading Go modules (go mod download) ---"
        exec go mod download
        ;;
    swag)
        echo "--- Generating Swagger documentation (swag init) ---"
        # Use -d to specify the directory to parse for Go files.
        exec swag init -g main.go -d ./src
        ;;
    init)
        echo "--- Installing Go development tools (air, swag) ---"
        # First, download modules to ensure tool versions are locked from go.mod
        go mod download
        # Install both tools in a single step
        go install github.com/air-verse/air github.com/swaggo/swag/cmd/swag
        echo "--- Tools installed successfully ---"
        ;;
    test)
        echo "--- Running Go tests ---"
        exec go test "$@"
        ;;
    *)
        echo "Error: Unknown dev command '$COMMAND'" >&2
        exit 1
        ;;
esac