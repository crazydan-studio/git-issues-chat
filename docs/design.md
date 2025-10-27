# Git Issues Chat - Design Documentation

## Overview
Git Issues Chat is a desktop application that transforms GitHub repositories into chat rooms, using Issues as topics and Comments as messages. This approach enables real-time collaboration and issue handling through a familiar chat interface.

## Architecture

### Technology Stack
- **Backend**: Golang for core application logic, data management, and GitHub API integration
- **Frontend**: Svelte with Tailwind CSS for the user interface
- **UI Framework**: Go-WebUI for desktop application deployment
- **Database**: SQLite for local data storage
- **Build System**: Vite for UI bundling

### Core Components

#### 1. Application Core (Golang)
- GitHub API integration for repository, issue, and comment management
- SQLite database for local caching and offline support
- WebUI integration for desktop application deployment
- Command-line interface for configuration and control

#### 2. User Interface (Svelte + Tailwind CSS)
- Chat-style interface for issue discussion
- Repository navigation and issue browsing
- Real-time message updates
- User authentication and preferences

#### 3. Data Flow
```
GitHub API ↔ Application Core ↔ SQLite Database ↔ UI Layer
```

## Project Structure
```
git-issues-chat/
├── build/          # Build scripts
├── deps/           # Third-party dependencies
├── dist/           # Build output
│   ├── bin/        # Executable files
│   ├── ui/         # UI build output
│   └── data/       # Application runtime data
├── docs/           # Documentation
│   └── stages/     # Stage development documentation
├── src/            # Source code
│   ├── main.go     # Application entry point
│   └── ui/         # Svelte UI project
└── VERSION         # Application version
```

## Development Phases

### Phase 0: Project Foundation (Completed)
- Basic project structure and build system
- Skeleton application with WebUI integration
- Simple welcome UI to verify technical approach

### Phase 1: GitHub Integration
- GitHub authentication and API access
- Repository listing and selection
- Issue and comment data retrieval

### Phase 2: Chat Interface Implementation
- Chat-style UI for issues and comments
- Real-time message updates
- Message composition and sending

### Phase 3: Advanced Features
- Offline support with SQLite caching
- Notification system
- User preferences and settings

### Phase 4: Finalization
- Performance optimization
- Testing and bug fixes
- Documentation and user guides

## Versioning
Application version is maintained in the `VERSION` file and automatically propagated to:
- `src/main.go` for the `--version` flag
- `src/ui/package.json` for UI build identification

## Build Process
The build process is handled by `build/build.sh` which:
1. Cleans previous build artifacts
2. Updates version information
3. Builds the Svelte UI with Tailwind CSS
4. Compiles the Golang application in both release and debug modes

## Runtime Requirements
- Supported browsers: Firefox, Chrome, Edge
- Environment variable: `WEBKIT_DISABLE_DMABUF_RENDERER=1` (for WebKit compatibility)
- Required directories: data storage and UI assets