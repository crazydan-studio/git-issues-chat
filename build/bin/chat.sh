#!/bin/bash

# Get the directory where this script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Set default values
BROWSER="firefox"
DEBUG=false

# Function to display help
show_help() {
  echo "Git Issues Chat - Desktop Application"
  echo ""
  echo "Usage: $0 [OPTIONS]"
  echo ""
  echo "Options:"
  echo "  --debug              Run the debug version of the application"
  echo "  --browser BROWSER    Specify browser to use (firefox, chrome, edge)"
  echo "  --help               Display this help message"
  echo ""
  echo "Examples:"
  echo "  $0                   Run the application with default settings"
  echo "  $0 --debug           Run the debug version"
  echo "  $0 --browser chrome  Run with Chrome browser"
  echo "  $0 --debug --browser edge  Run debug version with Edge browser"
}

# Parse command line arguments
while [[ $# -gt 0 ]]; do
  case $1 in
    --debug)
      DEBUG=true
      shift
      ;;
    --browser)
      BROWSER="$2"
      shift 2
      ;;
    --help)
      show_help
      exit 0
      ;;
    *)
      echo "Unknown option $1"
      echo "Use --help for usage information"
      exit 1
      ;;
  esac
done

# Set paths using relative paths from script location
DATA_PATH="$SCRIPT_DIR/data"
UI_PATH="$SCRIPT_DIR/ui"

# Automatically create data directory if it doesn't exist
if [ ! -d "$DATA_PATH" ]; then
  echo "Creating data directory: $DATA_PATH"
  mkdir -p "$DATA_PATH"
  if [ $? -ne 0 ]; then
    echo "Error: Failed to create data directory: $DATA_PATH"
    exit 1
  fi
fi

# Check if UI directory exists
if [ ! -d "$UI_PATH" ]; then
  echo "UI directory not found: $UI_PATH"
  echo "Required UI files are missing"
  exit 1
fi

# Check if executable exists
if [ "$DEBUG" = true ]; then
  EXECUTABLE="$SCRIPT_DIR/bin/chat-debug"
  echo "Running debug version..."
else
  EXECUTABLE="$SCRIPT_DIR/bin/chat"
  echo "Running release version..."
fi

if [ ! -f "$EXECUTABLE" ]; then
  echo "Executable not found: $EXECUTABLE"
  echo "Required executable is missing"
  exit 1
fi

# Permission checks
echo "Checking permissions..."

# Check if data directory is readable and writable
if [ ! -r "$DATA_PATH" ]; then
  echo "Error: Data directory is not readable: $DATA_PATH"
  exit 1
fi

if [ ! -w "$DATA_PATH" ]; then
  echo "Error: Data directory is not writable: $DATA_PATH"
  exit 1
fi

echo "Data directory permissions OK"

# Check if UI directory and its subdirectories are readable
if [ ! -r "$UI_PATH" ]; then
  echo "Error: UI directory is not readable: $UI_PATH"
  exit 1
fi

# Check if index.html exists in UI directory and is readable
if [ ! -r "$UI_PATH/index.html" ]; then
  echo "Error: UI index.html is not readable: $UI_PATH/index.html"
  exit 1
fi

# Check if assets directory exists and is readable
if [ -d "$UI_PATH/assets" ] && [ ! -r "$UI_PATH/assets" ]; then
  echo "Error: UI assets directory is not readable: $UI_PATH/assets"
  exit 1
fi

echo "UI directory permissions OK"

# Check if executable is executable
if [ ! -x "$EXECUTABLE" ]; then
  echo "Error: Executable is not executable: $EXECUTABLE"
  exit 1
fi

echo "Executable permissions OK"

# Set environment variable for WebKit
export WEBKIT_DISABLE_DMABUF_RENDERER=1

# Run the application
echo "Starting Git Issues Chat..."
echo "Data path: $DATA_PATH"
echo "UI path: $UI_PATH"
echo "Browser: $BROWSER"

"$EXECUTABLE" --data-path="$DATA_PATH" --ui-path="$UI_PATH" --browser="$BROWSER"