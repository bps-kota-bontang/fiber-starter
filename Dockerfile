# Stage 1: Build Stage
FROM golang:1.24-alpine AS builder

# Set environment variables for Go
ENV GO111MODULE=on \
    GOPROXY=https://proxy.golang.org

# Install build dependencies
RUN apk update && apk add --no-cache \
    git \
    bash \
    make \
    && rm -rf /var/cache/apk/*

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the entire application code
COPY . .

# Build the Go app (optimized for production)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o fiber-starter-service ./cmd

# Stage 2: Final Stage (Minimal, Production-Ready Image)
FROM alpine:latest

# Install SSL certificates to avoid certificate errors
RUN apk add --no-cache ca-certificates

# Add a non-root user for better security
RUN adduser -D -g '' user

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/fiber-starter-service .

# Use the non-root user to run the app
USER user

# Command to run the Go application
CMD ["./fiber-starter-service"]
