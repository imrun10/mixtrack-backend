# MixTrack Backend

A Go-based web service and CLI tool designed for managing music projects, featuring both a REST API server and command-line functionality for project initialization.

## Features

- **Web API**: Fiber-based HTTP server with health check endpoints
- **CLI Tool**: Command-line interface for managing music projects
- **Project Management**: Initialize and manage music project structures
- **Health Monitoring**: Built-in health check endpoint for service monitoring

## Prerequisites

- Go 1.24.1 or higher
- Git (for version control integration)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/imrun10/mixtrack-backend.git
cd mixtrack-backend
```

2. Install dependencies:

```bash
go mod download
```

3. Build the application:

```bash
go build -o mixtrack .
```

## Usage

### Web Server

Start the web server:

```bash
go run main.go
```

The server will start on port 8080 with the following endpoints:

- `GET /` - Welcome message
- `GET /health` - Health check endpoint

### CLI Tool

Use the CLI tool for project management:

```bash
./mixtrack --help
```

Initialize a new music project:

```bash
./mixtrack start <project-name>
```

## Project Structure

```text
mixtrack-backend/
├── main.go           # Web server entry point
├── go.mod           # Go module definition
├── go.sum           # Dependency checksums
├── cmd/             # CLI commands
│   └── root.go      # Root command definition
└── internal/        # Internal packages
    └── git/         # Git integration
        └── init.go  # Project initialization
```

## Dependencies

- [Fiber v2](https://github.com/gofiber/fiber) - Web framework
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [UUID](https://github.com/google/uuid) - UUID generation

## Development

### Running in Development

```bash
go run main.go
```

### Testing

```bash
go test ./...
```

### Building

```bash
go build -o mixtrack .
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | /        | Welcome message |
| GET    | /health  | Health check |
