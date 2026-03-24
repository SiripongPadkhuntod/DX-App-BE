# Dexter Transport Service

Building a robust backend infrastructure for Dexter Transport service using Go with Clean Architecture.

## 🚀 Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.21+)
- [Docker](https://www.docker.com/products/docker-desktop) & Docker Compose
- [Swag](https://github.com/swaggo/swag) (for API Documentation)

### Environment Setup

1. Copy the example environment file:
   ```bash
   cp config/local.env.example config/local.env
   ```
2. Update the values in `config/local.env` if necessary.

### Running with Docker

Start the PostgreSQL database:
```bash
docker-compose up -d
```

### Database Migrations

Run migrations using the Go CLI tool:

**Migrate Up:**
```bash
go run cmd/migrate/migrate.go -action=up
```

**Migrate Down:**
```bash
go run cmd/migrate/migrate.go -action=down
```

### Running the Server

Start the application:
```bash
go run cmd/server/main.go
```
The server will be available at `http://localhost:8080`.

## 📖 API Documentation

The project uses Swagger for documentation. After starting the server, you can access it at:
`http://localhost:8080/swagger/index.html`

To regenerate Swagger docs:
```bash
swag init -g cmd/server/main.go --parseDependency --parseInternal
```

## 🏗 Architecture

This project follows **Clean Architecture** (Port/Adapter pattern) to ensure scalability, testability, and maintainability.

- **`cmd/`**: Entry points for the application (Server, CLI tools).
- **`internal/app/`**: Core application logic.
  - **`domain/`**: Business entities and rules.
  - **`port/`**: Interfaces defining the boundaries (Service/Repository).
  - **`service/`**: Implementation of business logic.
  - **`handler/`**: HTTP handlers and request/response binding.
- **`internal/infrastructure/`**: Concrete implementations of external technologies (DB Client, etc.).
- **`internal/constant/`**: System-wide constants and error codes.
- **`pkg/v1/dto/`**: Shared Data Transfer Objects.

## 🛠 Tech Stack

- **Framework**: [Gin Gonic](https://gin-gonic.com/)
- **Database**: PostgreSQL
- **Migrations**: [golang-migrate](https://github.com/golang-migrate/migrate)
- **Documentation**: Swashbuckle (Swagger/OpenAPI)
- **Configuration**: [godotenv](https://github.com/joho/godotenv)
