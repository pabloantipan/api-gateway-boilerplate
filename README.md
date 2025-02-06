# REST API Gateway

A lightweight API Gateway implementation in Go that provides routing, authentication, and proxy capabilities for microservices.

## Features

- Firebase Authentication with email/password
- Reverse proxy to microservices
- Request/path routing
- Health check endpoint
- Path-based whitelisting
- Cloud logging integration

## Prerequisites

- Go 1.23 or higher
- Firebase project with service account
- Running microservices to proxy to

## Configuration

### Environment Variables (.env)

Just watch .env.example and see what is needed.

### Service Routes (config/routes.go)
This is just an example for this boiler plate.
```go
routes: map[string]Route{
    "/api/v1/players": {
        Path:      "/api/v1/players",
        TargetURL: "http://localhost:8085",
    },
}
```

## Installation

1. Clone the repository
```bash
git clone <repository-url>
```

2. Install dependencies
```bash
go mod download
```

3. Create .env file with required variables

4. Run the gateway
```bash
go run cmd/main.go
```

## API Endpoints

### Public Endpoints
- `GET /health` - Health check
- `POST /api/v1/auth/login` - Firebase authentication

### Protected Endpoints
All routes under `/api/v1/*` require Firebase authentication token

## Usage Examples

### Health Check
```bash
curl http://localhost:8080/health
```

### Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password"
  }'
```

### Protected Endpoint
```bash
curl http://localhost:8080/api/v1/players \
  -H "Authorization: Bearer <your-token>"
```

## Architecture

- **Presentation Layer**: HTTP handlers and middleware
- **Service Layer**: Business logic and routing
- **Infrastructure Layer**: Proxy implementation and external services

## Project Structure
```
├── cmd/
│   └── main.go
├── config/
│   └── routes.go
├── internal/
│   ├── presentation/
│   │   ├── handler/
│   │   └── middleware/
│   ├── service/
│   └── infrastructure/
│       └── proxy/
├── .env
└── go.mod
```

## Adding New Services

1. Add new route in `config/routes.go`:
```go
routes: map[string]Route{
    "/api/v1/new-service": {
        Path:      "/api/v1/new-service",
        TargetURL: "http://localhost:8086",
    },
}
```

2. Service will be automatically available through the gateway

## Error Handling

- 401 - Unauthorized (Invalid/missing token)
- 404 - Route not found
- 503 - Service unavailable

## Notes

- Trailing slashes are handled automatically
- Firebase tokens are validated on each protected request
- Proxy maintains original request method and headers