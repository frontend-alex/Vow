# Vow Server

Go HTTP API for Vow.

## Commands

```bash
go run ./cmd/api
go test ./...
make migrate-up
docker compose -f deployments/docker/docker-compose.yml up
```

The API listens on `:8080` by default and exposes `GET /health`.
