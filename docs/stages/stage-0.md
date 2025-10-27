# Stage 0: Project Initialization and Basic Setup

## Overview
This document describes the tasks completed in Stage 0 of the Git Issues Chat project. The goal of this stage was to establish the basic project structure, implement the foundational code, and verify that the technical approach is viable.

## Tasks Completed

### 1. Project Directory Structure Initialization
- Created the required directory structure:
  - `build`: Build scripts
  - `deps`: Third-party dependencies
  - `dist`: Build output
    - `bin`: Executable files
    - `ui`: UI build output
    - `data`: Application runtime data
  - `docs`: Documentation
    - `stages`: Stage development documentation
  - `src`: Source code
    - `ui`: UI project code

### 2. Version and License Setup
- Created `VERSION` file with version `0.1.0`
- Added Apache 2.0 license file

### 3. Golang Project Setup
- Initialized Go module with required dependencies:
  - `github.com/mattn/go-sqlite3` for SQLite database support
  - `github.com/webui-dev/go-webui/v2` for WebUI integration
- Created `src/main.go` with:
  - Command-line argument parsing for `--data-path`, `--ui-path`, `--browser`, and `--version`
  - SQLite database initialization function `initAppEnv`
  - WebUI window creation and browser selection based on `--browser` parameter
  - Support for Firefox, Chrome, and Edge browsers

### 4. Svelte + Tailwind CSS UI Setup
- Initialized Svelte project in `src/ui`
- Added Tailwind CSS support with proper configuration
- Created a welcome page with:
  - Full-screen centered layout
  - "👏 Hello, WebUI Desktop!" text with 64px font size
  - Proper viewport sizing
- Configured Vite build to output to `dist/ui`

### 5. Build System Implementation
- Created `build/build.sh` script that:
  - Reads version from `VERSION` file
  - Updates version in both Go source and UI package.json
  - Cleans `dist` directories
  - Builds UI with Tailwind CSS support
  - Clones go-webui as a dependency
  - Builds both release and debug versions of the Go application
  - Applies optimization flags (`-s -w`) to reduce executable size

### 6. Application Run Script
- Created `dist/run.sh` script that:
  - Handles `--debug` flag to run debug version
  - Accepts `--browser` parameter (default: firefox)
  - Sets required environment variable `WEBKIT_DISABLE_DMABUF_RENDERER=1`
  - Validates existence of required directories and executables
  - Passes correct parameters to the application

## Issues Encountered and Solutions

### 1. Tailwind CSS Build Issues
**Problem**: Initial Tailwind CSS setup failed with postcss plugin errors.
**Solution**: 
- Installed `@tailwindcss/postcss` package
- Updated `postcss.config.js` to use the correct plugin reference

### 2. Go-webui Integration Issues
**Problem**: Linker errors when building Go application due to missing webui library functions.
**Solution**:
- Cloned go-webui repository as a submodule in `deps/`
- Updated `go.mod` to replace the remote dependency with the local version
- Used appropriate build tags (`webui_std webui_no_cgo`) for compatibility

### 3. UI Build Warnings
**Problem**: Vite warned about `<script src="/webui.js">` not having `type="module"` attribute.
**Solution**: This is a non-blocking warning for our use case since WebUI handles this script differently than standard modules.

## Verification
The build system was successfully tested and produces:
- UI build output in `dist/ui`
- Release executable `dist/bin/chat`
- Debug executable `dist/bin/chat-debug`
- Functional run script that properly launches the application

## Next Steps
With the basic project structure and build system in place, subsequent stages will focus on implementing the core application functionality, including GitHub integration, issue tracking, and real-time messaging features.