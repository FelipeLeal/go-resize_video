#!/bin/sh

# Exit immediately if a command exits with a non-zero status.
set -e

# If the first argument is 'shell', drop into an interactive shell.
# This allows 'docker-compose run resize shell' or 'docker-compose up' (with command: shell)
# to provide an interactive session inside the container.
if [ "$1" = "shell" ]; then
    echo "---"
    echo "Entering interactive shell. Development commands are available via the 'dev' script."
    echo "Example: 'dev install-tools', then 'dev up' to start the server."
    echo "---"
    # We use /bin/bash for a better interactive experience, as it was installed in the Dockerfile.
    # The -i flag forces an interactive shell, which ensures a prompt is displayed with 'docker-compose up'.
    exec /bin/bash -i
fi

# If the command is not 'shell', just execute it.
# This allows running one-off commands like 'docker-compose run resize go version'.
exec "$@"
