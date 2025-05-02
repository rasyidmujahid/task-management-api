# Build stage
FROM golang:1.21-slim

WORKDIR /app

# Copy all files first
COPY . .

# Download dependencies
RUN go mod download && go build -o main

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