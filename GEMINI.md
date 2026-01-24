# Project: REST API Service — Ticket Booking

## Project Overview

This project is a Go-based RESTful API Service for Ticket Booking. It provides endpoints to manage events, availability, bookings, users and authentication. The codebase uses the Echo framework for routing and GORM as an ORM for interacting with a PostgreSQL database. The service includes a command-line interface for database migrations and seeding, and supports TLS for secure communication.

## Building and Running

### Prerequisites

- Go (version 1.20 or higher)
- PostgreSQL
- `air` (for auto-reloading during development)

### Installation

1.  Clone the repository.
2.  Install dependencies:

    ```bash
    go mod tidy
    ```

3.  Install development tools:

    ```bash
    make install-tools
    ```

### Configuration

1.  Copy the `.env.example` file to `.env`:

    ```bash
    cp .env.example .env
    ```

2.  Edit the `.env` file to configure your database connection, JWT secret, and other settings.

### Running the Application

-   **Development:**

    To run the application in development mode with auto-reloading, use the following command:

    ```bash
    make dev
    ```

    This will start the server and watch for file changes.

-   **Production:**

    To run the application in a production-like environment, you can build the binary and then execute it:

    ```bash
    go build -o myapp
    ./myapp
    ```

### Database Migrations

The project includes a simple command-line interface for managing database migrations.

-   **Run Migrations:**

    ```bash
    go run main.go migrate
    ```

-   **Rollback Migrations:**

    ```bash
    go run main.go rollback
    ```

-   **Fresh Migration with Seeding:**

    This command will roll back all migrations, run them again, and then seed the database with initial data.

    ```bash
    go run main.go migrate:fresh
    ```

## Development Conventions

### Code Style

The project follows standard Go conventions.

### Testing

There are no tests in the project. If you want to add tests, you should create files with the `_test.go` suffix.

### Contribution Guidelines

The `README.md` file provides detailed instructions for setting up and running the project. When contributing, please ensure that your changes are consistent with the existing code and that you update the documentation if necessary.

## Module Contexts (per-module)

The project maintains per-module context files under the `docs/` folder. See the following files for module-specific context:

- `docs/config.md` — configuration and DB initialization
- `docs/domain.md` — domain entities and interfaces (DDD)
- `docs/application.md` — application services (use-cases)
- `docs/dto.md` — request/response DTOs
- `docs/handler.md` — HTTP handlers
- `docs/infrastructure.md` — infrastructure implementations (persistence, etc.)
- `docs/models.md` — GORM models
- `docs/router.md` — routing
- `docs/seeder.md` — database seeding
- `docs/service.md` — legacy services (pre-DDD)
- `docs/tlsutil.md` — TLS helpers
- `docs/utils.md` — utility helpers

Use these files as the canonical per-module context for tools that read `GEMINI.md`.
