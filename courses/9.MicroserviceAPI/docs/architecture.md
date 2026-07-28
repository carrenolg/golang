# Architecture

## Overview

The service follows a strict layered architecture where each layer depends only on the layers below it. Dependency injection is wired manually in `app/app.go`.

```
┌─────────────────────────────────┐
│           app/                  │  HTTP — decode, call service, write response
├─────────────────────────────────┤
│          service/               │  Business logic, validation, orchestration
├─────────────────────────────────┤
│          domain/                │  Entities, repository interfaces, domain rules
├─────────────────────────────────┤
│      DB (MySQL via sqlx)        │  Persistence
└─────────────────────────────────┘
```

`dto/` carries data across the HTTP boundary (handler ↔ service). It is not a layer — it is a shared type package used by both `app/` and `service/`.

## Layers

### `app/` — Handlers

Handlers decode the HTTP request, call the service, and write the response. They do no business logic. All JSON responses go through the `writeResponse` helper in `handlers.go`.

Errors are type-asserted to `*errs.AppError` to extract the HTTP status code:

```go
appErr := err.(*errs.AppError)
writeResponse(w, appErr.Code, appErr)
```

`GET /customers/{id}` is the only endpoint with content negotiation: it returns XML when the request carries `Content-Type: application/xml`.

### `service/` — Business logic

Services receive and return DTOs, never domain entities. They enforce business rules before delegating persistence to the repository:

- `NewAccount`: validates amount ≥ 5000 and account type before saving.
- `MakeTransaction`: validates transaction type, then checks available balance for withdrawals before saving.

### `domain/` — Entities and repository interfaces

Domain entities model the database rows and carry domain-rule methods:

| Method | Entity | Rule |
|---|---|---|
| `CanWithdraw(amount)` | `Account` | Returns false if balance < amount |
| `IsWithdrawal()` | `Transaction` | True when type is `"withdrawal"` |
| `StatusAsText()` | `Customer` | Converts DB `"1"`/`"0"` to `"active"`/`"inactive"` |

Repository interfaces (`AccountRepository`, `CustomerRepository`) are defined here, next to the entities they operate on. This keeps the domain self-contained and testable without a database.

`CustomerRepositoryStub` is an in-memory implementation used for testing. It only implements `FindAll` — `ById` is not stubbed.

### `errs/` — Errors

`AppError` is the single error type used throughout the service. It carries an HTTP status code alongside the message, so handlers can set the correct response code without knowing HTTP details deeper in the stack.

```go
type AppError struct {
    Code    int
    Message string
}
```

Constructors: `NewNotFoundError` (404), `NewValidationError` (400), `NewUnexpectedError` (500).

### `logger/` — Logging

Thin wrapper around Uber's `zap` with production config and ISO8601 timestamps. Call `logger.Info`, `logger.Debug`, or `logger.Error` with optional `zap.Field` values.

## Database

Three active tables:

| Table | Description |
|---|---|
| `customers` | Customer profiles. `status`: `1` = active, `0` = inactive. |
| `accounts` | Bank accounts linked to customers. `account_type`: `saving` or `checking`. |
| `transactions` | Deposit/withdrawal history linked to accounts. |

Two additional tables exist in the schema for future use: `users` (authentication) and `refresh_token_store` (JWT refresh tokens). Neither is wired up yet.

`sqlx` is used for all DB access. Struct fields **must** have `db:"column_name"` tags for row scanning to work.

### `SaveTransaction` — SQL transaction

`AccountRepositoryDb.SaveTransaction` wraps two operations in a single SQL transaction:

1. `INSERT` into `transactions`
2. `UPDATE accounts SET amount = amount ± ?`

On error, the transaction is rolled back. On success, it must be committed before reading the updated balance.

## Request flow example

`POST /customers/2000/accounts/95470` with `{"amount": 100, "transaction_type": "withdrawal"}`:

```
MakeTransaction (handler)
  └─ decode body → TransactionRequest
  └─ set AccountId, CustomerId from URL vars
  └─ service.MakeTransaction(request)
        └─ request.Validate()              ← type must be deposit/withdrawal, amount ≥ 0
        └─ repo.FindById(accountId)        ← load current balance
        └─ account.CanWithdraw(100)        ← reject if balance < 100
        └─ repo.SaveTransaction(t)         ← INSERT + UPDATE in one SQL tx
        └─ return TransactionResponse      ← includes new balance
  └─ writeResponse(200, response)
```
