# Stage 0: Project Initialization and Basic Setup

## Overview
This document describes the tasks completed in Stage 0 of the Git Issues Chat project. The goal of this stage was to establish the basic project structure, implement the foundational code, and verify that the technical approach is viable.

## Tasks Completed

### 1. Project Directory Structure Initialization
- Created the required directory structure:
  - `build`: Build scripts
    - `bin`: Runtime scripts
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
  - "👏 Hello, WebUI Desktop!" text with 72px font size
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
  - Copies `build/bin/chat.sh` to `dist/chat.sh`

### 6. Application Run Script
- Created `build/bin/chat.sh` script that:
  - Handles `--debug` flag to run debug version
  - Accepts `--browser` parameter (default: firefox)
  - Supports `--help` for usage information
  - Automatically creates data directory if it doesn't exist
  - Sets required environment variable `WEBKIT_DISABLE_DMABUF_RENDERER=1`
  - Validates existence and permissions of required directories and executables
  - Passes correct parameters to the application
- Script is copied to `dist/chat.sh` during build process

## Issues Encountered and Solutions

### 1. Tailwind CSS Build Issues
**Problem**: Initial Tailwind CSS setup failed with postcss plugin errors.
**Solution**: 
- Installed `@tailwindcss/postcss` package
- Updated `postcss.config.js` to use the correct plugin reference

### 2. Tailwind CSS Font Size Issue
**Problem**: Tailwind CSS predefined font size classes (like `text-8xl` for 72px) were not working in the built application. The generated CSS was missing these classes due to Tailwind's content-based purging in v4, which only includes classes it can detect in the source files.

**Solution**: 
- Changed from predefined classes like `text-8xl` to arbitrary value notation `text-[72px]`
- Arbitrary values in Tailwind CSS v4+ are more reliably included in the build
- This approach explicitly tells Tailwind to generate a CSS rule for exactly 72px font size
- Updated `src/ui/src/App.svelte` to use `class="text-[72px] font-bold"`

**Verification**: 
- After rebuilding, the generated CSS includes `.text-\[72px\]{font-size:72px}`
- The UI now properly displays the text at 72px font size as requested

### 3. Go-webui Integration Issues
**Problem**: Linker errors when building Go application due to missing webui library functions.
**Solution**:
- Cloned go-webui repository as a submodule in `deps/`
- Updated `go.mod` to replace the remote dependency with the local version
- Used appropriate build tags (`webui_std webui_no_cgo`) for compatibility

### 4. UI Build Warnings
**Problem**: Vite warned about `<script src="/webui.js">` not having `type="module"` attribute.
**Solution**: This is a non-blocking warning for our use case since WebUI handles this script differently than standard modules.

### 5. File Organization Changes
**Problem**: The run script was initially located in `dist/run.sh`, which was not ideal for the project structure.
**Solution**:
- Moved the run script to `build/bin/chat.sh`
- Updated the build script to copy `build/bin/chat.sh` to `dist/chat.sh` after building
- Updated the script to use paths relative to its own location for better portability

### 6. Permission and Directory Creation
**Problem**: The application was automatically creating the data directory, which wasn't the desired behavior.
**Solution**:
- Removed automatic data directory creation from `src/main.go`'s `initAppEnv` function
- Moved the responsibility to the run script (`build/bin/chat.sh`)
- Updated the script to automatically create the data directory if it doesn't exist
- Added comprehensive permission checks for all required files and directories

## Verification
The build system was successfully tested and produces:
- UI build output in `dist/ui`
- Release executable `dist/bin/chat`
- Debug executable `dist/bin/chat-debug`
- Functional run script `dist/chat.sh` that properly launches the application
- Script supports `--help` for usage information and proper error handling

## Next Steps
With the basic project structure and build system in place, subsequent stages will focus on implementing the core application functionality, including GitHub integration, issue tracking, and real-time messaging features.