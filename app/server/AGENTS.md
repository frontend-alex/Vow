# Harness App Server Agent Instructions

## Core rules

These rules apply whenever I make changes to the `app/server` Go codebase.

1. I will always wrap code examples in fenced code blocks.
2. I will always update `internal/docs/openapi.go` whenever I change, add, remove, or rename a server API endpoint, request DTO, response DTO, route, status code, or documented API behavior.
3. I will not leave hardcoded values directly inside business logic, handlers, repositories, or routers. If a value must be configurable, I will add it to `internal/config/config.go`, reference it from the environment, and document it in `.env.example`.
4. Before creating new logic, I will scan the codebase to check whether an existing function, type, middleware, repository method, DTO, error, or helper can be reused or extended.
5. I will follow the existing feature-based architecture and Domain-Driven Design structure.
6. I will keep handlers thin. Handlers sanitize input, validate input, call the service, and return standardized JSON responses. They must not contain business logic.
7. I will keep business rules inside the service layer.
8. I will keep direct database access inside the repository layer.
9. I will use existing shared packages where possible, especially shared response, request, validation, config, and error utilities.
10. I will keep naming consistent with the existing codebase.

## External Source Inspection with opensrc

When I need to understand an external dependency, framework, SDK, package, or GitHub repository, I will inspect the real source code with `opensrc` before making implementation decisions.

I will not guess how a dependency behaves when the source code can be inspected.

### When I will use opensrc

I will use `opensrc` when:

- I need to understand how an installed library works.
- I need to understand how a GitHub repository implements a feature.
- I need to inspect validation behavior, middleware behavior, handlers, helpers, types, or exported APIs.
- I am unsure about the correct usage of a package.
- I need to verify the actual implementation instead of relying only on documentation or assumptions.

### Where opensrc stores source code

When `opensrc` fetches source code, it creates a local cache folder named `.opensrc`.

The cache can look like this:

```code
/.opensrc/
├── repos/
│   └── github.com/
│       ├── colinhacks/
│       │   └── zod/
│       │       └── v3.24.1/
│       └── vercel/
│           └── ai/
│               └── main/
└── sources.json
```

## Project architecture

This is a Go server codebase using feature-based architecture.

When I create a new feature, I will place it under:

```code
server/internal/[feature-name]
```

Example:

```code
server/internal/morningroutine
```

Inside each feature folder, I will separate responsibilities into clear files and layers.

Expected feature structure:

```code
server/internal/[feature-name]/
├── handler.go
├── router.go
├── dto.go
├── service.go
└── repo.go
```

If the feature grows, I can split files further while keeping the same layer boundaries.

Example of a larger feature:

```code
server/internal/auth/
├── handler.go
├── login.go
├── logout.go
├── register.go
├── router.go
├── dto.go
├── service.go
└── repo.go
```

## Handler rules

Handlers live in:

```code
server/internal/[feature-name]/handler.go
```

If a handler file becomes crowded, I will split controller actions into separate files.

Example:

```code
server/internal/auth/login.go
server/internal/auth/logout.go
server/internal/auth/register.go
```

I will keep a main `handler.go` file for the shared handler struct, constructor, and dependencies.

### Handler responsibilities

A handler is responsible for:

1. Reading request input.
2. Sanitizing request input.
3. Validating request input.
4. Calling the correct service method.
5. Returning the service result using the shared response helpers.

A handler is not responsible for:

1. Business rules.
2. Database queries.
3. Complex branching logic.
4. Duplicate validation that belongs in DTO tags or services.
5. Creating feature-specific error behavior that belongs in the service layer.

### Handler response rules

Every handler response must use the response helpers from:

```code
internal/shared/response/json.go
```

Depending on the controller behavior, I will return the correct helper, such as:

```code
Ok
Created
OkNoMessage
CreatedNoMessage
```

I will not manually write inconsistent JSON response shapes inside handlers.

### Handler validation and sanitization

In `handler.go`, I will sanitize and validate request input before calling the service.

I will use the shared sanitization middleware/helper from:

```code
internal/shared/request/sanitize.go
```

I will define validation rules on DTO structs using:

```code
github.com/go-playground/validator/v10
```

Validation tags belong in `dto.go`.

Example request DTO:

```code
type AuthUserRequest struct {
	ID                  int64  `json:"id" validate:"required,gt=0"`
	Email               string `json:"email" validate:"required,email,max=255" sanitize:"trim"`
	Name                string `json:"name" validate:"omitempty,max=255" sanitize:"trim"`
	OnboardingCompleted bool   `json:"onboardingCompleted"`
}
```

## Router rules

Feature routes live in:

```code
server/internal/[feature-name]/router.go
```

Inside `router.go`, I will register all API endpoints for that feature.

Route format:

```code
/v1/api/[feature-name]/[action-or-resource]
```

Example:

```code
func AuthenticationRoutes(mux *http.ServeMux, handler Handler) {
	mux.HandleFunc("POST /v1/api/auth/login", handler.Login)
	mux.HandleFunc("POST /v1/api/auth/register", handler.Register)
	mux.HandleFunc("POST /v1/api/auth/logout", handler.Logout)
}
```

After creating the feature route registration function, I will wire the feature inside:

```code
server/internal/routes/[feature-name].go
```

Example:

```code
func Authentication(mux *http.ServeMux, deps Dependencies) {
	repository := auth.NewRepository(deps.DB)
	service := auth.NewService(repository, deps.Config.JWTSecret)
	handler := auth.NewHandler(service)

	auth.AuthenticationRoutes(mux, handler)
}
```

Then I will register the feature builder inside:

```code
server/internal/routes/routes.go
```

Example:

```code
func Router(mux *http.ServeMux, deps Dependencies) {
	Authentication(mux, deps)
	Onboarding(mux, deps)
}
```

## DTO rules

DTOs live in:

```code
server/internal/[feature-name]/dto.go
```

I will create DTOs for request and response payloads used by handlers and services.

### DTO naming rules

I will name DTOs using the feature name plus `Request` or `Response`.

Examples:

```code
AuthRequest
AuthResponse
LoginRequest
LoginResponse
RegisterRequest
RegisterResponse
```

For feature-specific actions, I will use clear action names.

Example:

```code
MorningRoutineCreateRequest
MorningRoutineCreateResponse
MorningRoutineUpdateRequest
MorningRoutineUpdateResponse
```

### Request DTO rules

Request DTOs must include validation tags when validation is required.

Bad example:

```code
type AuthUserRequest struct {
	ID                  int64  `json:"id"`
	Email               string `json:"email"`
	Name                string `json:"name"`
	OnboardingCompleted bool   `json:"onboardingCompleted"`
}
```

Good example:

```code
type AuthUserRequest struct {
	ID                  int64  `json:"id" validate:"required,gt=0"`
	Email               string `json:"email" validate:"required,email,max=255" sanitize:"trim"`
	Name                string `json:"name" validate:"omitempty,max=255" sanitize:"trim"`
	OnboardingCompleted bool   `json:"onboardingCompleted"`
}
```

### Response DTO rules

Response DTOs define the shape of the `data` field returned by the shared API response wrapper.

The shared API response shape is:

```code
type APIResponse struct {
	Success      bool        `json:"success"`
	Message      *string     `json:"message"`
	ErrorMessage *string     `json:"error_message"`
	ErrorStatus  *int        `json:"error_status"`
	ErrorCode    *string     `json:"error_code"`
	UserMessage  *string     `json:"user_message"`
	Data         interface{} `json:"data"`
}
```

A feature response DTO should describe what goes inside `Data`.

Example:

```code
type AuthResponse struct {
	UserID int64  `json:"userId"`
	Email  string `json:"email"`
	Token  string `json:"token,omitempty"`
}
```

## Service rules

Services live in:

```code
server/internal/[feature-name]/service.go
```

The service layer owns the business logic.

A service is responsible for:

1. Business rules.
2. Feature-specific checks.
3. Calling repository methods.
4. Deciding which shared error to return.
5. Coordinating multiple repository calls when needed.
6. Returning response DTOs to the handler.

A service must not:

1. Read raw HTTP request bodies.
2. Write HTTP responses.
3. Register routes.
4. Directly depend on `http.ResponseWriter` or `*http.Request` unless the existing codebase already has a clear pattern for it.
5. Perform raw database access that belongs in the repository.

Example service check:

```code
_, err := s.repository.GetUserByEmail(ctx, email)
if err == nil {
	return AuthResponse{}, sharederrors.AuthErrors.EmailAlreadyTaken
}
```

When a service detects a known failure case, I will return the correct error object from:

```code
internal/shared/errors/errors.go
```

If no matching shared error exists, I will create a new error in the shared errors file.

Example:

```code
ExampleError: APIError{
	ErrorCode:   "EXX_002",
	StatusCode:  http.StatusNotImplemented,
	Message:     "Not implemented.",
	UserMessage: "This feature is not available yet.",
},
```

I will keep error codes consistent with the existing naming and numbering pattern.

## Repository rules

Repositories live in:

```code
server/internal/[feature-name]/repo.go
```

The repository layer is responsible for direct database access using GORM.

A repository is responsible for:

1. Creating database records.
2. Reading database records.
3. Updating database records.
4. Deleting database records.
5. Returning database errors to the service.

A repository must not contain:

1. Complex business rules.
2. HTTP response logic.
3. Request validation logic.
4. Feature orchestration logic.

Only database-specific checks should exist in the repository. Business validation must happen in the service layer.

## Config and environment rules

When I need a configurable value, I will add it to:

```code
internal/config/config.go
```

I will also add the matching environment variable to:

```code
.env
.env.example
```

I will avoid scattering constants across the codebase.

Examples of values that should go through config:

```code
JWT secrets
Token expiry durations
External service URLs
API keys
Timeouts
Feature flags
Rate limits
Upload limits
```

## OpenAPI documentation rules

Whenever I change the server API, I will update:

```code
internal/docs/openapi.go
```

This applies when I:

1. Add an endpoint.
2. Remove an endpoint.
3. Rename an endpoint.
4. Change request payloads.
5. Change response payloads.
6. Change status codes.
7. Change authentication requirements.
8. Change error responses.

I will keep the OpenAPI documentation aligned with the actual route, DTO, and response behavior.

## Reuse-first rule

Before adding new code, I will inspect the existing codebase for reusable patterns.

I will check for existing:

```code
Handlers
Services
Repositories
DTOs
Errors
Middleware
Response helpers
Request helpers
Config patterns
Validation patterns
Route registration patterns
```

If similar logic already exists, I will reuse it or extend it instead of duplicating it.

## Naming and style rules

I will use clear, consistent names.

I will avoid typos in file names, function names, struct names, and route names.

Correct examples:

```code
service.go
repo.go
handler.go
router.go
AuthenticationRoutes
NewRepository
NewService
NewHandler
```

Incorrect examples:

```code
serviceре.go
Repo.go
contoller.go
interna/routes/routes.go
scheama
codabase
```

I will follow Go formatting standards and run formatting where appropriate.

Expected formatting tools:

```code
gofmt
go test ./...
```

When possible, I will run tests after changes. If I cannot run tests, I will clearly state that tests were not run.

## Implementation checklist

Before I finish a server change, I will verify the following:

1. I placed new feature code under `server/internal/[feature-name]`.
2. I kept handler, service, repository, router, and DTO responsibilities separate.
3. I sanitized and validated handler input.
4. I used validation tags in request DTOs.
5. I returned standardized responses from `internal/shared/response/json.go`.
6. I kept business logic in `service.go`.
7. I kept database access in `repo.go`.
8. I reused existing functions and patterns where possible.
9. I added or reused shared errors from `internal/shared/errors/errors.go`.
10. I moved configurable values into `internal/config/config.go` and environment files.
11. I registered feature routes in the feature router, feature builder, and main router.
12. I updated `internal/docs/openapi.go` for API changes.
13. I formatted the Go code.
14. I ran tests when possible.
15. I reported any skipped tests, assumptions, or follow-up work clearly.
