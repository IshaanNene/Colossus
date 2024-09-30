#!/bin/bash

# Build script for ARM Emulator (Colossus)

set -e # Exit immediately if a command exits with a non-zero status

# Define variables
BUILD_DIR="./bin"
SRC_DIR="./cmd"

# Create the build directory if it doesn't exist
mkdir -p "$BUILD_DIR"

# Build the emulator
echo "Building the ARM emulator..."
go build -o "$BUILD_DIR/colossus" "$SRC_DIR/main.go"

echo "Build completed successfully!"
