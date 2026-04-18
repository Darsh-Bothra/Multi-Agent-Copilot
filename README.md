# Multi-Agent Copilot

Backend API for transaction tracking and (via the database schema) group expense splitting. The service is a Go application using the [Gin](https://github.com/gin-gonic/gin) web framework and PostgreSQL.

## Prerequisites

- [Go](https://go.dev/dl/) 1.26 or newer (see `apps/api/go.mod`)
- [PostgreSQL](https://www.postgresql.org/) 14+ recommended

## Project layout

| Path | Purpose |
|------|---------|
| `apps/api/cmd/server` | Application entrypoint |
| `apps/api/handlers` | HTTP handlers |
| `apps/api/service` | Business logic |
| `apps/api/repository` | Database access |
| `apps/api/internal/config` | Environment-based configuration |
| `apps/api/internal/db` | PostgreSQL connection and SQL migrations |

## Configuration

Configuration is loaded from environment variables. Optionally create `apps/api/.env` (see `.gitignore`); if `.env` is missing, values fall back to defaults or the system environment.

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | HTTP listen port |
| `DB_HOST` | `localhost` | PostgreSQL host |
| `DB_PORT` | `5432` | PostgreSQL port |
| `DB_USER` | `postgres` | Database user |
| `DB_PASSWORD` | `password` | Database password |
| `DB_NAME` | `upi_app` | Database name |
| `DB_SSLMODE` | `disable` | PostgreSQL `sslmode` query parameter |

## Database setup

1. Create a database (name should match `DB_NAME`, e.g. `upi_app`).

2. Apply migrations in chronological order using your preferred method (`psql`, a migration tool, etc.). Migration files live under `apps/api/internal/db/migrations/`:

   - `20260415172225_create_transactions` — creates the `transactions` table and enables `uuid-ossp`.
   - `20260417170710_create_group_expense` — creates `users`, `groups`, `group_members`, `expenses`, and `splits`.

Example with `psql` (adjust connection string as needed):

```bash
psql "postgres://postgres:password@localhost:5432/upi_app?sslmode=disable" \
  -f apps/api/internal/db/migrations/20260415172225_create_transactions.up.sql

psql "postgres://postgres:password@localhost:5432/upi_app?sslmode=disable" \
  -f apps/api/internal/db/migrations/20260417170710_create_group_expense.up.sql
```

## Run the API

From the `apps/api` directory:

```bash
cd apps/api
go run ./cmd/server
```

The server listens on `PORT` (default `8080`).

## HTTP API

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/ping` | Health check; returns `{ "message": "pong" }` |
| `GET` | `/transactions` | List transactions |
| `POST` | `/transactions` | Create a transaction |

**Create transaction** — JSON body:

```json
{
  "amount": 42.5,
  "merchant": "Coffee Shop"
}
```

Successful creation responds with `201 Created` and a confirmation message.

## Development

```bash
cd apps/api
go test ./...
go build -o bin/server ./cmd/server
```

## License

Add a license file at the repository root if you plan to distribute this project.
