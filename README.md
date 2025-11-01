# Git Issues Chat

A desktop chat application that transforms GitHub repositories into chat rooms using Git Issues as topics and comments as messages.

## Overview

Git Issues Chat provides a familiar chat interface for GitHub collaboration, making it easier to track issues, discuss solutions, and coordinate development efforts. By treating repositories as chat rooms, issues as topics, and comments as messages, the application streamlines communication and issue management.

## Features

- **GitHub Integration**: Direct integration with GitHub repositories
- **Chat Interface**: Familiar messaging interface for issue discussions
- **Multi-Browser Support**: Works with Firefox, Chrome, and Edge
- **Offline Capability**: Local SQLite database for offline access
- **Real-time Updates**: Automatic synchronization with GitHub
- **Cross-Platform**: Runs on Linux, Windows, and macOS
- **User Management**: Authentication, profile management, and password changes
- **Action Logging**: Track user activities and actions
- **App Locking**: Secure application with password protection

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
│   ├── handlers/   # Backend handler functions
│   ├── main.go     # Application entry point
│   └── ui/         # Svelte UI project
└── VERSION         # Application version
```

## Development Setup

### Prerequisites

- Go 1.21 or later
- Node.js 20+ and npm
- Git

### Building the Project

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd git-issues-chat
   ```

2. Run the build script:
   ```bash
   ./build/build.sh
   ```

This will:
- Build the Svelte UI with Tailwind CSS
- Compile the Go application
- Generate both release and debug versions
- Copy the runtime script to the dist directory

### Running the Application

After building, run the application using the provided script:

```bash
./dist/chat.sh [--debug] [--browser firefox|chrome|edge] [--help]
```

Options:
- `--debug`: Run the debug version of the application
- `--browser`: Specify the browser to use (default: firefox)
- `--help`: Display usage information

### Development Commands

#### UI Development

Navigate to the UI directory:
```bash
cd src/ui
```

Install dependencies:
```bash
npm install
```

Run development server:
```bash
npm run dev
```

Build UI for production:
```bash
npm run build
```

#### Go Development

Build the Go application:
```bash
go build -o dist/bin/chat src/main.go
```

Run with parameters:
```bash
./dist/bin/chat --data-path=./dist/data --ui-path=./dist/ui --browser=firefox
```

## Application Components

### Backend (Golang)
- **Handlers**: Modular handler functions for different application features
  - `app.go`: Application-level functions (user management, info, etc.)
  - `platform.go`: Git platform management
  - `repo.go`: Repository management
  - `issue.go`: Issue and comment management

### Frontend (Svelte + Tailwind CSS)
- **Component Organization**:
  - `app/`: Application-level components (authentication, main layout, dialogs)
  - `chat/`: Chat panel components
  - `issue/`: Issue-related components (lists, panels, comments)
  - `repo/`: Repository-related components (platforms, repositories)
  - `user/`: User-related components (profile, password management)
- **State Management**: Svelte stores for application state
- **Bridge Layer**: JavaScript functions for communicating with backend

## Versioning

The application version is maintained in the `VERSION` file and is automatically applied to both the Go application and UI package during the build process.

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.

## Contributing

See the development documentation in the `docs/` directory for information about project phases and implementation details.