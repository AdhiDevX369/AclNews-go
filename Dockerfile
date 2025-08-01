# Multi-stage build for optimized image size
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Install git (required for some Go modules)
RUN apk add --no-cache git ca-certificates

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o anime-news-ai cmd/app/main.go

# Final stage - minimal image
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN addgroup -g 1001 app && adduser -D -s /bin/sh -u 1001 -G app app

# Set working directory
WORKDIR /home/app

# Copy binary from builder stage
COPY --from=builder /app/anime-news-ai .

# Change ownership to app user
RUN chown -R app:app /home/app

# Switch to non-root user
USER app

# Expose port (optional, as this is a CLI app)
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD ./anime-news-ai --health || exit 1

# Command to run
CMD ["./anime-news-ai"]
