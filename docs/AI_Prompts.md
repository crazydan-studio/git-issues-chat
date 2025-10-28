# Git Issues Chat - AI Implementation Prompts

This document contains the complete, organized prompts and requirements for implementing the Git Issues Chat desktop application. These prompts are designed to be executed by an AI assistant to implement the project systematically.

## Project Overview

Design and implement a desktop chat application named `Git Issues Chat` that uses GitHub repositories as chat platforms, repositories as chat rooms, Issues as topics, and Comments as messages. This enables real-time issue handling and collaboration through a chat interface.

## Technology Stack

- **Backend**: Golang for core application logic, data management, and GitHub API integration
- **Frontend**: Svelte with Tailwind CSS for the user interface
- **UI Framework**: Go-WebUI for desktop application deployment
- **Database**: SQLite for local data storage
- **Build System**: Vite for UI bundling

## Project Structure

```
git-issues-chat/
├── build/          # Build scripts
│   ├── bin/        # Runtime scripts
│   └── build.sh    # Main build script
├── deps/           # Third-party dependencies
├── dist/           # Build output
│   ├── bin/        # Executable files
│   ├── ui/         # UI build output
│   ├── data/       # Application runtime data
│   └── chat.sh     # Runtime script
├── docs/           # Documentation
│   └── stages/     # Stage development documentation
├── src/            # Source code
│   ├── main.go     # Application entry point
│   └── ui/         # Svelte UI project
└── VERSION         # Application version
```

## Phase 1: Project Foundation (Stage 0)

### Task 1: Initialize Project Structure

Create the project directory structure as specified above with the following requirements:

1. Initialize git repository
2. Create VERSION file with version "0.1.0"
3. Add Apache 2.0 LICENSE file
4. Set up proper .gitignore to exclude dist/ and deps/ directories

### Task 2: Set up Golang Project

Create `src/main.go` with the following functionality:

1. Command-line argument parsing:
   - `--data-path`: Path to data directory (required)
   - `--ui-path`: Path to UI directory (required)
   - `--browser`: Browser to use (firefox, chrome, edge) (required)
   - `--version`: Print version information ("Git Issues Chat v0.1.0")

2. Database initialization function `initAppEnv`:
   - Takes dataPath parameter
   - Creates data directory if it doesn't exist
   - Initializes SQLite database at `{dataPath}/chat.db`
   - Returns database object

3. WebUI integration:
   - Create WebUI window
   - Set root folder to ui-path
   - Show browser window with index.html using specified browser
   - Wait for window to close

4. Dependencies:
   - github.com/mattn/go-sqlite3 for SQLite support
   - github.com/webui-dev/go-webui/v2 for WebUI integration

### Task 3: Set up Svelte + Tailwind CSS UI

Initialize Svelte project in `src/ui` with:

1. Tailwind CSS integration:
   - Install tailwindcss, postcss, autoprefixer
   - Configure tailwind.config.js and postcss.config.js
   - Add Tailwind directives to app.css

2. Welcome page requirements:
   - Full-screen centered layout
   - Text: "👏 Hello, WebUI Desktop!" with 72px font size
   - Body should occupy full viewport

3. Build configuration:
   - Set output directory to `../../dist/ui`
   - Include webui.js script in index.html

4. Package configuration:
   - Set package name to "git-issues-chat-ui"
   - Set version to match project version

### Task 4: Implement Build System

Create `build/build.sh` script that:

1. Reads version from VERSION file
2. Updates version in:
   - src/main.go (for --version flag)
   - src/ui/package.json
3. Cleans dist directories
4. Builds UI with Tailwind CSS support
5. Fetches go-webui dependency:
   - Clone to deps/go-webui if not exists
   - Update if exists
6. Builds Go application:
   - Release version: dist/bin/chat with optimization flags (-s -w)
   - Debug version: dist/bin/chat-debug with debug flags
7. Copies build/bin/chat.sh to dist/chat.sh
8. Includes error checking for all critical steps

### Task 5: Implement Runtime Script

Create `build/bin/chat.sh` script that:

1. Command-line arguments:
   - `--debug`: Run debug version
   - `--browser`: Specify browser (firefox, chrome, edge)
   - `--help`: Display usage information

2. Path management:
   - Uses paths relative to script location
   - DATA_PATH: {script_dir}/data
   - UI_PATH: {script_dir}/ui
   - EXECUTABLE: {script_dir}/bin/chat or chat-debug

3. Directory management:
   - Automatically creates data directory if it doesn't exist
   - Validates UI directory exists

4. Permission checks:
   - Data directory: readable and writable
   - UI directory and files: readable
   - Executable: executable

5. Environment:
   - Sets WEBKIT_DISABLE_DMABUF_RENDERER=1
   - Passes correct parameters to application

### Task 6: Documentation

Create/update documentation files:

1. `docs/stages/stage-0.md`:
   - Document all tasks completed
   - Record issues encountered and solutions
   - Include technical details about Tailwind CSS font size issue:
     * Problem: Predefined classes not working due to Tailwind v4 purging
     * Solution: Use arbitrary value notation (text-[72px])
   - Document file organization changes
   - Document permission and directory creation changes

2. `docs/design.md`:
   - Update project structure diagram
   - Document technical notes about Tailwind CSS handling
   - Update runtime script information

3. `README.md`:
   - Update project structure
   - Update usage instructions
   - Document prerequisites (Node.js 20+)

### Task 7: Version Control

Commit all changes with appropriate commit messages including:
- Project structure initialization
- Core application implementation
- UI implementation
- Build system implementation
- Documentation updates
- Error handling improvements

## Technical Requirements and Constraints

### Tailwind CSS Font Size Handling

Due to Tailwind CSS v4's content-based purging, predefined font size classes may not be included in the build. For reliable font size control, use arbitrary value notation (e.g., `text-[72px]`) instead of predefined classes (e.g., `text-8xl`).

### Data Directory Management

The application should not automatically create the data directory. This responsibility should be moved to the runtime script, which should automatically create the data directory if it doesn't exist, providing better separation of concerns.

### Error Handling

All critical operations should include proper error handling:
- Go-webui fetching failures
- Application executable build failures
- Browser window opening failures
- File permission and existence checks

### Path Management

All scripts should use paths relative to their own location for better portability rather than absolute paths or complex path resolution logic.

## Quality Assurance

1. All builds should complete without errors
2. Generated executables should run correctly
3. UI should display properly with correct font sizes
4. All documentation should be accurate and up-to-date
5. Git commits should be atomic and well-documented
6. Error messages should be clear and helpful
7. Help documentation should be comprehensive

## Verification Steps

1. Run build script successfully
2. Verify dist directory contents:
   - bin/chat and bin/chat-debug executables
   - ui/ directory with built assets
   - chat.sh script
   - data/ directory
3. Test runtime script with various options
4. Verify help functionality works
5. Check error handling for missing files/directories
6. Confirm all documentation is accurate and complete