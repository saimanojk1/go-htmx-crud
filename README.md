# Go HTMX CRUD

A simple task management web application built with Go and HTMX for learning purposes.

## Features

- Create, read, update, and delete tasks
- Real-time UI updates with HTMX
- PostgreSQL database integration
- Clean architecture with service layers

## Prerequisites

- Go 1.24+
- Docker
- PostgreSQL client tools (`psql`, `pg_isready`)

## Quick Start

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd go-htmx-crud
   ```

2. **Start the database**
   ```bash
   ./local-db-startup.sh
   ```

3. **Install dependencies**
   ```bash
   go mod tidy
   ```

4. **Run the application**
   ```bash
   go run cmd/main.go
   ```

5. **Open your browser**
   Navigate to `http://localhost:4000`

## Project Structure

```
├── cmd/
│   ├── main.go          # Application entry point
│   └── api/
│       └── api.go       # HTTP server setup
├── services/
│   ├── database/        # Database connection and operations
│   └── tasks/           # Page routes & operations
├── templates/           # HTML templates
├── schema.sql           # Database schema
└── local-db-startup.sh  # Database setup script
```

## Database

The application uses PostgreSQL with a simple `tasks` table:
- `id` (SERIAL PRIMARY KEY)
- `task` (VARCHAR(200))
- `done` (BOOLEAN)

The database runs on port 15432 to avoid conflicts with existing PostgreSQL installations.

## Development

- The application runs on `localhost:4000`
- Templates are automatically parsed from the `templates/` directory
- Database connection is managed through the `database` service layer

