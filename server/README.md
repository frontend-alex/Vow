# Server

Go API server for Vow.

## Layout

- `cmd/api/main.go` starts the HTTP API process.
- `internal/app` wires the Gin router.
- `internal/auth` contains authentication packages used only by this server.
- `internal/shared/validation` validates DTOs with go-playground/validator v10.

## Development

```bash
go run ./cmd/api
go test ./...
go vet ./...
go build ./...
```

The initial API exposes `GET /health`.
