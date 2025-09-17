#!/bin/bash
set -e

echo "🔨 Building Feast Friends API..."

# Create bin directory if it doesn't exist
mkdir -p bin

# Build for Linux (production)
echo "🐧 Building for Linux..."
GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/feast-friends-api-linux cmd/server/main.go

# Build for current OS (development)
echo "💻 Building for development..."
go build -o bin/feast-friends-api cmd/server/main.go

echo "✅ Build completed successfully!"
ls -la bin/