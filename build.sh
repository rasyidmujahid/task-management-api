#!/bin/bash
set -e

echo "Building application..."
go mod tidy
CGO_ENABLED=0 GOOS=linux go build -o main .

echo "Build complete. Current directory contents:"
ls -la

echo "Running application..."
./main 