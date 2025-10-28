#!/bin/bash

# Get the directory where this script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

# Read version from VERSION file
VERSION=$(cat "$PROJECT_ROOT/VERSION")

echo "Building Git Issues Chat v$VERSION..."

# Clean dist directories
echo "Cleaning dist directories..."
rm -rf "$PROJECT_ROOT/dist/bin" "$PROJECT_ROOT/dist/ui"
mkdir -p "$PROJECT_ROOT/dist/bin" "$PROJECT_ROOT/dist/ui" "$PROJECT_ROOT/dist/data"

# Update version in Go file
echo "Updating version in Go source..."
sed -i "s/Git Issues Chat v[0-9]*\.[0-9]*\.[0-9]*/Git Issues Chat v$VERSION/g" "$PROJECT_ROOT/src/main.go"

# Update version in UI package.json
echo "Updating version in UI package.json..."
cd "$PROJECT_ROOT/src/ui"
npm version "$VERSION" --no-git-tag-version --force

# Build UI
echo "Building UI..."
cd "$PROJECT_ROOT/src/ui"
npm run build

# Check if UI build was successful
if [ $? -ne 0 ]; then
    echo "Error: UI build failed"
    exit 1
fi

# Build Go application
echo "Building Go application..."
cd "$PROJECT_ROOT"

# Get go-webui as submodule
echo "Getting go-webui dependency..."
mkdir -p deps
if [ ! -d "deps/go-webui" ]; then
    git clone https://github.com/webui-dev/go-webui.git deps/go-webui
    # Check if git clone was successful
    if [ $? -ne 0 ]; then
        echo "Error: Failed to clone go-webui repository"
        exit 1
    fi
else
    cd deps/go-webui && git pull && cd "$PROJECT_ROOT"
    # Check if git pull was successful
    if [ $? -ne 0 ]; then
        echo "Error: Failed to update go-webui repository"
        exit 1
    fi
fi

# Build release version
echo "Building release version..."
go build -tags "webui_std webui_no_cgo" -ldflags "-s -w" -o dist/bin/chat src/main.go

# Check if release build was successful
if [ $? -ne 0 ]; then
    echo "Error: Failed to build release version"
    exit 1
fi

# Build debug version
echo "Building debug version..."
go build -tags "webui_log webui_std webui_no_cgo" -o dist/bin/chat-debug src/main.go

# Check if debug build was successful
if [ $? -ne 0 ]; then
    echo "Error: Failed to build debug version"
    exit 1
fi

# Copy chat.sh script to dist directory
echo "Copying chat.sh script to dist directory..."
cp "$PROJECT_ROOT/build/bin/chat.sh" "$PROJECT_ROOT/dist/"

# Check if copy was successful
if [ $? -ne 0 ]; then
    echo "Error: Failed to copy chat.sh script to dist directory"
    exit 1
fi

echo "Build completed successfully!"