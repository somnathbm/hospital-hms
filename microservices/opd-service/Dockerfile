# --- Build stage ---
FROM golang:1.23-alpine AS build

# Setup working directory
WORKDIR /app

# Install required dependencies to build binary
RUN apk update && apk add --no-cache git

# Copy go mod and go sum file
COPY go.mod go.sum ./
RUN go mod download

# Copy entire source code
COPY . .

# Build the binary
RUN go build -o opd-service ./cmd/server

# --- Run stage ---
FROM alpine:latest AS server

# Create a non-root user for safety
RUN useradd -D -g '' opduser

# Setup working directory
WORKDIR /app

# Copy the binary file from build stage
COPY --from=build /app/opd-service .

# Change the ownership
RUN chown -R opduser:opduser /app

# Run as opduser
USER opduser

# Expose gRPC and HTTP port
EXPOSE 50051 8080

# Run the binary
ENTRYPOINT ["./opd-service"]