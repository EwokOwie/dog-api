# Animal Photos API

A backend service to retrieve animal breed photos. Currently supports dogs via the [Dog CEO API](https://dog.ceo/dog-api/), with an extensible architecture for adding other animals.

## Features

- RESTful API with JSON responses
- Polymorphic animal interface for extensibility
- Swagger UI documentation
- Structured logging with slog
- Graceful shutdown
- Table-driven tests with mock server

## Requirements

- Go 1.21+

## Getting Started

```bash
# Clone the repository
git clone https://github.com/EwokOwie/dog-api.git
cd dog-api

# Install dependencies
go mod download

# Run the server
go run ./cmd/web

# Or with custom port
go run ./cmd/web -addr=:3000
```

The server starts at http://localhost:8080

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/animals` | List available animals |
| GET | `/api/v1/animals/{animal}/breeds` | List breeds for an animal |
| GET | `/api/v1/animals/{animal}/breeds/{breed}/photo` | Get a random photo for a breed |
| GET | `/health` | Health check |
| GET | `/docs/` | Swagger UI documentation |
| GET | `/api/openapi.yaml` | OpenAPI specification |

## Example Requests

```bash
# List available animals
curl http://localhost:8080/api/v1/animals

# List dog breeds
curl http://localhost:8080/api/v1/animals/dog/breeds

# Get a husky photo
curl http://localhost:8080/api/v1/animals/dog/breeds/husky/photo

# Health check
curl http://localhost:8080/health
```

## Running Tests

```bash
# Run all tests
go test ./...

# Verbose output
go test ./... -v

# With coverage
go test ./... -cover
```

## Project Structure

```
.
├── api/
│   └── openapi.yaml        # OpenAPI specification
├── cmd/
│   └── web/
│       ├── main.go         # Application entry point
│       ├── routes.go       # Route definitions
│       ├── handlers.go     # HTTP handlers
│       ├── helpers.go      # Response helpers
│       └── middleware.go   # HTTP middleware
├── internal/
│   ├── assert/             # Test assertion helpers
│   └── models/
│       ├── animal.go       # Animal interface and service
│       ├── dog.go          # Dog implementation
│       └── errors.go       # Error types
├── go.mod
└── README.md
```

