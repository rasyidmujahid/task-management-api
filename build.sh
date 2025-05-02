#!/bin/bash
set -e

echo "Current directory: $(pwd)"
echo "Directory contents:"
ls -la

echo "Building application..."
go mod tidy
CGO_ENABLED=0 GOOS=linux go build -o /app/main .

echo "Build complete. Checking /app directory:"
ls -la /app

echo "Running application..."
cd /app && ./main 