#!/bin/bash

# This command ensures that the script will exit immediately if any command fails.
# It's a good practice for shell scripting to avoid unexpected behavior.
set -e

# Define the URL to test.
# The '${1:-"..."}' syntax means: use the first argument passed to the script ($1),
# but if no argument is provided, use "http://localhost:8080/health" as the default.
# This makes the script flexible for different environments.
TARGET_URL=${1:-"http://localhost:8080/health"}

# Set the maximum number of times the script should try to connect.
MAX_RETRIES=5
# Set the time (in seconds) to wait between retries.
RETRY_DELAY=5

echo "Running smoke test against: $TARGET_URL"

# A 'for' loop that will run from 1 up to the MAX_RETRIES value.
for ((i=1; i<=MAX_RETRIES; i++)); do
    # Use 'curl' to make an HTTP request to the target URL.
    # -s: Silent mode. Don't show progress meter or error messages.
    # -o /dev/null: Discard the response body. We only care about the status code.
    # -w "%{http_code}": Print only the HTTP status code to standard output.
    # We store this status code in the HTTP_STATUS variable.
    HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" $TARGET_URL)

    # Check if the HTTP status code is 200 (OK).
    if [ "$HTTP_STATUS" -eq 200 ]; then
        echo "✅ Smoke test passed! Received status 200."
        # Exit the script with a success code (0).
        exit 0
    else
        # If the status is not 200, print a failure message.
        echo "Attempt $i/$MAX_RETRIES failed with status $HTTP_STATUS. Retrying in $RETRY_DELAY seconds..."
        # Wait for the specified delay before the next attempt.
        sleep $RETRY_DELAY
    fi
done

# If the loop finishes without a successful connection, it means all retries failed.
echo "❌ Smoke test failed after $MAX_RETRIES attempts."
# Exit the script with a failure code (1).
exit 1