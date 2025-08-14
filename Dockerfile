# Use the ffmpeg base image directly
FROM jrottenberg/ffmpeg:6.1-alpine

# Set working directory
WORKDIR /app

# WORKAROUND: Temporarily switch to HTTP repositories to bypass SSL/TLS issues.
# Then, update the package list and install all required packages.
RUN sed -i 's/https/http/' /etc/apk/repositories && \
    apk update && \
    apk add --no-cache ca-certificates go entr

# Copy the dev script and make it executable
COPY src/dev.sh /usr/local/bin/dev.sh
RUN chmod +x /usr/local/bin/dev.sh

# PATCH: Modify dev.sh to run entr in non-interactive mode for Docker.
RUN sed -i "s/entr -r/entr -n -r/" /usr/local/bin/dev.sh

# Ensure the mounted dev.sh will have execute permissions
RUN chmod +x /usr/local/bin/dev.sh

# Expose the port the app runs on
EXPOSE 8080

# Set the entrypoint to run our development script
ENTRYPOINT ["/bin/sh", "-c", "chmod +x /usr/local/bin/dev.sh && exec /usr/local/bin/dev.sh"]