FROM jrottenberg/ffmpeg:6.1-alpine

WORKDIR /app

# WORKAROUND: Temporarily switch to HTTP repositories to bypass SSL/TLS issues.
# Then, update the package list and install all required packages.
RUN sed -i 's/https/http/' /etc/apk/repositories && \
    apk update && \
    apk add --no-cache ca-certificates go entr

COPY src/dev.sh /usr/local/bin/dev.sh
# Fix line endings (CRLF -> LF) and make the script executable
RUN sed -i 's/\r$//' /usr/local/bin/dev.sh && chmod +x /usr/local/bin/dev.sh

# Expose the port the app runs on
EXPOSE 8080