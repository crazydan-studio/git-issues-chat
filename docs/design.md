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
│   ├── bin/        # Runtime scripts
│   └── build.sh    # Main build script
├── deps/           # Third-party dependencies
├── dist/           # Build output
│   ├── bin/        # Executable files
│   ├── ui/         # UI build output
│   ├── data/       # Application runtime data
│   └── chat.sh     # Copied runtime script
├── docs/           # Documentation
│   └── stages/     # Stage development documentation
├── src/            # Source code
│   ├── handlers/   # Backend handler functions
│   │   └── params.go # Strongly-typed parameter structures
│   ├── main.go     # Application entry point
│   ├── types/      # Backend type definitions
│   └── ui/         # Svelte UI project
└── VERSION         # Application version
```

## Development Phases

### Phase 0: Project Foundation (Completed)
- Basic project structure and build system
- Skeleton application with WebUI integration
- Simple welcome UI to verify technical approach
- Implementation of runtime script with proper error handling and help support

### Phase 1: GitHub Integration
- GitHub authentication and API access
- Repository listing and selection
- Issue and comment data retrieval

### Phase 2: Chat Interface Implementation (Completed)
- Complete chat-style UI for issues and comments
- Platform, repository, and issue management
- User authentication and profile management
- Action logging and app locking features
- Real-time message updates and composition

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
5. Copies `build/bin/chat.sh` to `dist/chat.sh`

## Runtime Script
The application is run using `dist/chat.sh` which:
- Supports `--debug` flag to run debug version
- Accepts `--browser` parameter (default: firefox)
- Provides `--help` for usage information
- Automatically creates data directory if needed
- Performs comprehensive permission checks
- Uses paths relative to its own location for portability

## Runtime Requirements
- Supported browsers: Firefox, Chrome, Edge
- Environment variable: `WEBKIT_DISABLE_DMABUF_RENDERER=1` (for WebKit compatibility)
- Required directories: data storage and UI assets

## Technical Notes

### Tailwind CSS Font Size Handling
Due to Tailwind CSS v4's content-based purging, predefined font size classes may not be included in the build. For reliable font size control, arbitrary value notation (e.g., `text-[72px]`) is used instead of predefined classes (e.g., `text-8xl`).

### Data Directory Management
The application no longer automatically creates the data directory. This responsibility has been moved to the runtime script, which automatically creates the data directory if it doesn't exist, providing better separation of concerns.

### Component Architecture
The frontend components are organized by functionality:
- `app/`: Application-level components (authentication, main layout, dialogs)
- `chat/`: Chat panel components
- `issue/`: Issue-related components (lists, panels, comments)
- `repo/`: Repository-related components (platforms, repositories)
- `user/`: User-related components (profile, password management)

### State Management
Application state is managed through strongly-typed Svelte stores with clear separation of concerns:
- Application state (user info, selected items)
- UI state (dialog visibility, notifications)
- Data lists (platforms, repositories, issues, comments, participants)
- Action logs

All stores are now strongly typed with TypeScript interfaces defined in `src/ui/src/lib/types.ts`.

### Bridge Layer
Communication between frontend and backend is handled through a TypeScript bridge layer that:
- Provides generic function calling mechanism for Go handlers
- Handles parameter serialization and return value deserialization
- Implements result-based error handling instead of exception interception
- Exposes specific wrapper functions for all backend handlers with strong typing
- Uses strongly-typed parameter structures following the "ParamsOf + FunctionName" naming convention