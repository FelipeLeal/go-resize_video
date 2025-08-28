FROM golang:1.24-alpine

# Set the working directory early.
WORKDIR /app

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

# Copy module and tools files. The entrypoint script will use them to run Go commands.
COPY go.mod go.sum ./
COPY tools ./tools/

# Copy the dev script (our task runner) and the entrypoint script.
COPY --chmod=755 .docker/dev.sh /usr/local/bin/dev
COPY --chmod=755 .docker/entrypoint.sh /usr/local/bin/entrypoint.sh

# Set the entrypoint for the container.
ENTRYPOINT ["entrypoint.sh"]
