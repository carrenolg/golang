# Banking Microservice

A RESTful banking API built in Go as part of a microservices course. It manages customers, bank accounts, and deposit/withdrawal transactions backed by MySQL.

## Prerequisites

- Go 1.24+
- MySQL 8+

## Setup

**1. Create the database and load the schema:**

```bash
mysql -u root -p < docs/tables.sql
```

This creates the `banking` database with seed data for customers and accounts.

**2. Configure environment variables:**

Create `bank/.env`:

```env
DB_DATASOURCE=root:password@tcp(localhost:3306)/banking?parseTime=true
SERVER_ADDRESS=
SERVER_PORT=8080
```

`SERVER_ADDRESS` can be left empty to bind to all interfaces on the given port.

## Running

```bash
cd bank
go run main.go
```

## API

| Method | Path | Description |
|---|---|---|
| `GET` | `/customers?status=active\|inactive` | List customers |
| `GET` | `/customers/{customer_id}` | Get one customer |
| `POST` | `/customers/{customer_id}/account` | Create a bank account |
| `POST` | `/customers/{customer_id}/accounts/{account_id}` | Make a transaction |

- `GET /customers/{customer_id}` supports content negotiation: send `Content-Type: application/xml` to receive an XML response.
- Account creation requires `amount >= 5000` and `account_type` of `saving` or `checking`.
- Transactions accept `transaction_type` of `deposit` or `withdrawal`. Withdrawals are rejected if the balance is insufficient.

See `docs/api.http` for ready-to-run sample requests and `docs/openapi.yaml` for the full API specification.

## Project structure

```
bank/
├── app/        # HTTP handlers and router wiring
├── service/    # Business logic and validation
├── domain/     # Entities, repository interfaces, domain methods
├── dto/        # Request/response types
├── errs/       # AppError with HTTP status codes
└── logger/     # Zap wrapper
```

See `docs/architecture.md` for a detailed explanation of the design.
