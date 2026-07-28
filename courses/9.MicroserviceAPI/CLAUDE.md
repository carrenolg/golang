# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

All commands run from the `bank/` directory:

```bash
# Run the server
go run main.go

# Build
go build ./...

# Test
go test ./...

# Vet
go vet ./...
```

## Environment

The server reads from `.env` (loaded via `godotenv`) or system environment variables:

| Variable | Description | Default |
|---|---|---|
| `DB_DATASOURCE` | MySQL DSN (required) | — |
| `SERVER_ADDRESS` | Bind address | `:8080` |
| `SERVER_PORT` | Port | `8080` |

The MySQL database is named `banking`. Schema and seed data are in `docs/tables.sql`.

## Architecture

The service uses a strict layered architecture with dependency injection wired in `app/app.go`:

```
app/          HTTP handlers — decode requests, call services, write JSON responses
service/      Business logic — validation, orchestration, domain rules
domain/       Entities + repository interfaces + domain methods
dto/          Request/response types that cross the HTTP boundary
errs/         AppError with HTTP status codes (400/404/500)
logger/       Zap wrapper (Info/Debug/Error)
```

**Data flow:** Handler → Service (uses `dto`) → Repository interface (uses `domain`) → DB implementation

**Dependency rule:** outer layers depend on inner layers; `domain` has no knowledge of `service` or `app`.

### Key patterns

- Repository interfaces live in `domain/` alongside the entities they operate on (`AccountRepository`, `CustomerRepository`).
- DB implementations (`accountRepositoryDb`, `customerRepositoryDb`) satisfy those interfaces using `sqlx`.
- `customerRepositoryStub.go` provides an in-memory stub for testing.
- DTOs have `Validate()` methods; domain entities have business-rule methods (e.g., `CanWithdraw`, `IsWithdrawal`, `StatusAsText`).
- All errors are returned as `*errs.AppError`, which implements `error` and carries an HTTP status code. Handlers type-assert to `*errs.AppError` to get the code.
- `SaveTransaction` in `accountRepositoryDb` wraps an INSERT + UPDATE in a SQL transaction and returns the updated balance on the `Amount` field.

## API endpoints

| Method | Path | Description |
|---|---|---|
| GET | `/customers?status=active\|inactive` | List customers |
| GET | `/customers/{customer_id}` | Get one customer |
| POST | `/customers/{customer_id}/account` | Create account (min amount 5000, type: saving/checking) |
| POST | `/customers/{customer_id}/accounts/{account_id}` | Make transaction (deposit/withdrawal) |

Sample requests are in `docs/api.http`.
