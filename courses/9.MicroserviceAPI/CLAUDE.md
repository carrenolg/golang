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

# Run a single test
go test ./... -run TestName

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
- DB implementations (`AccountRepositoryDb`, `customerRepositoryDb`) satisfy those interfaces using `sqlx`.
- All `domain` structs scanned from DB rows **must** have `db:"column_name"` struct tags — `sqlx` uses them for mapping.
- `CustomerRepositoryStub` is an in-memory test double that only implements `FindAll` (not `ById`).
- DTOs have `Validate()` methods; domain entities have business-rule methods (e.g., `CanWithdraw`, `IsWithdrawal`, `StatusAsText`).
- All errors are returned as `*errs.AppError`, which implements `error` and carries an HTTP status code. Handlers type-assert to `*errs.AppError` to get the code.
- `writeResponse` in `app/handlers.go` is the single helper for writing JSON responses across all handlers.
- `SaveTransaction` in `accountRepositoryDb` wraps an INSERT + UPDATE in a SQL transaction and returns the updated balance on the `Amount` field.
- DB stores account/customer status as `"1"`/`"0"`; the API returns `"active"`/`"inactive"` via `StatusAsText()`.

### Content negotiation

`GET /customers/{customer_id}` supports both JSON (default) and XML. Send `Content-Type: application/xml` to get an XML response.

## API endpoints

| Method | Path | Description |
|---|---|---|
| GET | `/customers?status=active\|inactive` | List customers |
| GET | `/customers/{customer_id}` | Get one customer (JSON or XML) |
| POST | `/customers/{customer_id}/account` | Create account (min amount 5000, type: saving/checking) |
| POST | `/customers/{customer_id}/accounts/{account_id}` | Make transaction (deposit/withdrawal) |

Sample requests are in `docs/api.http`.
