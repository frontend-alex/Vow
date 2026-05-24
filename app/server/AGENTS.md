# Backend Feature Rule

When creating or changing backend features in this Go server, always check and follow these rules:

- Inspect the existing package patterns before editing. Follow the current handler, service, repository, DTO, route, middleware, config, response, and errors structure.
- If creating or changing a route that returns a response, add every new API error in `internal/shared/errors/errors.go` only. Each error must include `ErrorCode`, `StatusCode`, `Message`, and `UserMessage`, and must be added to `ErrorMessages`.
- If creating or changing a route that returns a response, render errors only through `internal/shared/response/json.go`, usually with `response.AppError` or `response.AppErrorWithMessage`.
- Follow strict direct DDD boundaries. The handler must not contain core business logic, database logic, or feature logic.
- Handlers should only decode and validate requests, read route/context data, call the service, and write the HTTP response.
- Services own the feature/business flow and must return the appropriate centralized errors from `internal/shared/errors/errors.go` so handlers can render them through `internal/shared/response/json.go`.
- Keep database access in repositories. Services coordinate business flow and transactions, but must not scatter query details across handlers.
- When creating a new request DTO, always add validation tags/schema and validate it at the boundary with `request.DecodeAndValidate`.
- Do not duplicate request DTO validation in services. Only put checks in services when the rule is real domain/business logic that cannot be expressed by request validation.
- When adding or using an environment variable, define and load it in `internal/config/config.go`. Application code must read configuration from the config object only, not directly from `os.Getenv`.
- Keep Swagger/OpenAPI docs up to date with all added, changed, or removed routes, request bodies, responses, status codes, and response shapes.
- Prefer the smallest correct implementation. Do not add backward compatibility, extra abstractions, or helper layers without a concrete need.
- Before finishing, run `gofmt` on changed Go files and `go test ./...` from `app/server`.
