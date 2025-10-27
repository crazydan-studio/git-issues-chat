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
else
    cd deps/go-webui && git pull && cd "$PROJECT_ROOT"
fi

# Build release version
echo "Building release version..."
go build -tags "webui_std webui_no_cgo" -ldflags "-s -w" -o dist/bin/chat src/main.go

# Build debug version
echo "Building debug version..."
go build -tags "webui_log webui_std webui_no_cgo" -o dist/bin/chat-debug src/main.go

echo "Build completed successfully!"