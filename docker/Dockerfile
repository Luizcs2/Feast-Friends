<<<<<<< Updated upstream
# TODO - Luiz
# 1. Use official Go image as base
# 2. Set working directory
# 3. Copy go.mod and go.sum first (for better caching)
# 4. Download dependencies
# 5. Copy source code
# 6. Build the application
# 7. Use minimal image for final stage (alpine)
# 8. Copy binary and set entrypoint
# 9. Expose port 8080
# 10. Set up non-root user for security
=======
# ---- Build Stage (Stage 1) ----
# This first stage, named "builder", contains the Go compiler and tools.
# It's like a temporary workshop for building our application binary.
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container.
WORKDIR /app

# Copy dependency files first to leverage Docker's build cache.
# If these files don't change, Docker won't re-download all dependencies on every build.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application's source code into the builder stage.
COPY . .

# Compile the Go application into a single, static binary.
# CGO_ENABLED=0: Disables Cgo, which is needed to create a static binary that can run in a minimal container.
# GOOS=linux: Ensures the binary is compiled for a Linux OS (which Docker containers use).
# -o /main: Specifies the output path and name of the compiled application.
# ./cmd/api/: This should be the path to your main package. **Remember to change this if yours is different.**
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /main ./cmd/api/


# ---- Final Stage (Stage 2) ----
# This stage builds the final image. It starts from a minimal base image
# that does NOT contain any of the Go tools or source code.
FROM alpine:latest

# Set the working directory for the final container.
WORKDIR /root/

# The key step: copy ONLY the compiled binary from the "builder" stage.
# This makes the final image incredibly small and secure because it contains
# nothing but our application and its bare-minimum OS dependencies.
COPY --from=builder /main .

# Expose the port that the application listens on.
EXPOSE 8080

# The command to run when the container starts. This executes our application binary.
CMD ["./main"]
>>>>>>> Stashed changes
