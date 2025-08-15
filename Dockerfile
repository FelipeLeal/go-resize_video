FROM golang:1.24-alpine AS builder

# Set the working directory inside the container.
WORKDIR /app

# Copy the application source code.
COPY . .

# Install entr, which is a key part of your dev script.
RUN apk update && \
    apk add --no-cache entr

# Install Air
RUN go install github.com/air-verse/air@v1.62.0

# Build the main Go application.
# This will create a compiled binary.
RUN go mod tidy
RUN go build -o main .

# --- Final stage to run the application ---
# Use a smaller, cleaner base image to run the final application.
FROM jrottenberg/ffmpeg:6.1-alpine

# Set the working directory inside the container.
WORKDIR /app

# Copy the built binary from the builder stage.
COPY --from=builder /app/main .

# Expose the port the app runs on.
EXPOSE 8080

# This is the new command. Air will now handle the hot-reloading.
CMD ["./main"]