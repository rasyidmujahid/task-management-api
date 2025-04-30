# Build stage
FROM golang:1.21-bullseye AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM debian:bullseye-slim

WORKDIR /app

# Install ca-certificates for HTTPS
RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary from builder
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./main"] 