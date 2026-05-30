# Graph Report - .  (2026-05-30)

## Corpus Check
- Corpus is ~36,580 words - fits in a single context window. You may not need a graph.

## Summary
- 629 nodes · 1256 edges · 31 communities (26 shown, 5 thin omitted)
- Extraction: 83% EXTRACTED · 17% INFERRED · 0% AMBIGUOUS · INFERRED: 217 edges (avg confidence: 0.8)
- Token cost: 0 input · 0 output

## Community Hubs (Navigation)
- [[_COMMUNITY_Basic Auth Tests|Basic Auth Tests]]
- [[_COMMUNITY_Memory State Storage|Memory State Storage]]
- [[_COMMUNITY_API Router Auth|API Router Auth]]
- [[_COMMUNITY_Session Tests|Session Tests]]
- [[_COMMUNITY_Authenticator Tests|Authenticator Tests]]
- [[_COMMUNITY_Middleware Utilities|Middleware Utilities]]
- [[_COMMUNITY_WebAuthn Credentials|WebAuthn Credentials]]
- [[_COMMUNITY_Audit Event Tests|Audit Event Tests]]
- [[_COMMUNITY_OIDC Providers|OIDC Providers]]
- [[_COMMUNITY_Error Handling|Error Handling]]
- [[_COMMUNITY_Standard Logging|Standard Logging]]
- [[_COMMUNITY_Token Tests|Token Tests]]
- [[_COMMUNITY_Server CI Docs|Server CI Docs]]
- [[_COMMUNITY_Token Manager|Token Manager]]
- [[_COMMUNITY_Session Manager|Session Manager]]
- [[_COMMUNITY_OIDC Client|OIDC Client]]
- [[_COMMUNITY_TOTP Manager|TOTP Manager]]
- [[_COMMUNITY_Noop Auditor|Noop Auditor]]
- [[_COMMUNITY_Basic Auth Middleware|Basic Auth Middleware]]
- [[_COMMUNITY_JWT Middleware|JWT Middleware]]
- [[_COMMUNITY_Session Middleware|Session Middleware]]
- [[_COMMUNITY_Storage Interfaces|Storage Interfaces]]
- [[_COMMUNITY_Graphify Project Guide|Graphify Project Guide]]
- [[_COMMUNITY_Graphify Query Tools|Graphify Query Tools]]
- [[_COMMUNITY_OpenCode Config|OpenCode Config]]
- [[_COMMUNITY_Package Dependencies|Package Dependencies]]
- [[_COMMUNITY_API Entry Point|API Entry Point]]
- [[_COMMUNITY_Gin Application|Gin Application]]
- [[_COMMUNITY_Validation Library|Validation Library]]

## God Nodes (most connected - your core abstractions)
1. `NewInMemoryUserStore()` - 52 edges
2. `NewInMemoryCredentialStore()` - 38 edges
3. `Context` - 34 edges
4. `New()` - 27 edges
5. `Authenticator` - 23 edges
6. `NewTokenManager()` - 23 edges
7. `InMemoryCredentialStore` - 22 edges
8. `T` - 22 edges
9. `T` - 21 edges
10. `NewInMemorySessionStore()` - 21 edges

## Surprising Connections (you probably didn't know these)
- `Development Commands` --semantically_similar_to--> `go vet ./...`  [INFERRED] [semantically similar]
  server/README.md → .github/workflows/server-ci.yml
- `Development Commands` --semantically_similar_to--> `go test -race ./...`  [INFERRED] [semantically similar]
  server/README.md → .github/workflows/server-ci.yml
- `Development Commands` --semantically_similar_to--> `go build ./...`  [INFERRED] [semantically similar]
  server/README.md → .github/workflows/server-ci.yml
- `server/** Paths` --references--> `Server`  [INFERRED]
  .github/workflows/server-ci.yml → server/README.md
- `NewRouter()` --calls--> `Logger`  [INFERRED]
  server/internal/app/router.go → server/internal/auth/audit/stdlib.go

## Hyperedges (group relationships)
- **graphify Knowledge Graph Components** — agents_graphify_out, agents_god_nodes, agents_community_structure, agents_cross_file_relationships [EXTRACTED 1.00]
- **server-ci Go Quality Pipeline** — server_ci_gofmt_check, server_ci_go_vet, server_ci_go_test_race, server_ci_go_build [EXTRACTED 1.00]
- **Server Runtime Layout** — readme_cmd_api_main, readme_internal_app, readme_gin_router, readme_http_api_process [EXTRACTED 1.00]

## Communities (31 total, 5 thin omitted)

### Community 0 - "Basic Auth Tests"
Cohesion: 0.08
Nodes (36): BasicAuthWrapper, MockAuditor, MockFailingAuditor, SessionManagerWrapper, SourceExtractor, TokenManagerWrapper, NewBasicAuthWrapper(), NewSessionManagerWrapper() (+28 more)

### Community 1 - "Memory State Storage"
Cohesion: 0.07
Nodes (23): OIDCState, RWMutex, Context, Duration, SessionData, Time, User, WebAuthnCredential (+15 more)

### Community 2 - "API Router Auth"
Cohesion: 0.10
Nodes (21): env(), main(), NewRouter(), Authenticator, generateID(), GenerateResetToken(), generateVerificationToken(), NewAuthenticator() (+13 more)

### Community 3 - "Session Tests"
Cohesion: 0.11
Nodes (37): T, T, TestGenerateSessionID(), TestManager_Concurrency(), TestManager_Create(), TestManager_Delete(), TestManager_Get(), TestManager_Refresh() (+29 more)

### Community 4 - "Authenticator Tests"
Cohesion: 0.15
Nodes (37): contains(), containsInside(), TestAuthenticator_Authenticate(), TestAuthenticator_BcryptCost(), TestAuthenticator_ChangePassword(), TestAuthenticator_ConcurrentRegistrations(), TestAuthenticator_DisableTOTPWithInvalidCode(), TestAuthenticator_EmailVerification_SSO() (+29 more)

### Community 5 - "Middleware Utilities"
Cohesion: 0.09
Nodes (29): ContextKey, CookieExtractor, CookieWriter, ErrorHandler, HeaderExtractor, DefaultErrorHandler(), GetSessionID(), GetUserID() (+21 more)

### Community 6 - "WebAuthn Credentials"
Cohesion: 0.10
Nodes (22): AuthenticatorTransport, Credential, CredentialAssertion, CredentialCreation, ParsedCredentialAssertionData, ParsedCredentialCreationData, Context, CredentialStore (+14 more)

### Community 7 - "Audit Event Tests"
Cohesion: 0.12
Nodes (24): Actor, Actor, contains(), redactEmail(), redactIPAddress(), redactString(), TestAuditEvent_Structure(), TestEventResultConstants() (+16 more)

### Community 8 - "OIDC Providers"
Cohesion: 0.11
Nodes (16): Config, IDTokenVerifier, Provider, NewAppleProvider(), BaseOIDCProvider, NewGoogleProvider(), OAuth2Provider, NewOAuth2Provider() (+8 more)

### Community 9 - "Error Handling"
Cohesion: 0.12
Nodes (21): Error, Definition, Error, As(), Internal(), WithCause(), WithExtra(), WithUserMessage() (+13 more)

### Community 10 - "Standard Logging"
Cohesion: 0.16
Nodes (21): DefaultStdLogger(), NewStdLogger(), ProductionStdLogger(), BenchmarkStdLogger_Log(), BenchmarkStdLogger_LogWithRedaction(), TestDefaultStdLogger(), TestProductionStdLogger(), TestStdLogger_ComplexEvent() (+13 more)

### Community 11 - "Token Tests"
Cohesion: 0.25
Nodes (22): NewTokenManager(), ParseUnverified(), TestGenerateTokenID(), TestNewTokenManager(), TestParseUnverified(), TestTokenManager_ConcurrentGeneration(), TestTokenManager_GenerateAccessToken(), TestTokenManager_GenerateTokenPair() (+14 more)

### Community 12 - "Server CI Docs"
Cohesion: 0.10
Nodes (22): Development Commands, Go API Server, GET /health, internal/auth, Server, Vow, server-ci, actions/checkout@v6 (+14 more)

### Community 13 - "Token Manager"
Cohesion: 0.19
Nodes (14): Claims, Config, generateTokenID(), TokenManager, TokenPair, TokenType, RegisteredClaims, Context (+6 more)

### Community 14 - "Session Manager"
Cohesion: 0.16
Nodes (11): Context, Duration, SessionData, Config, CreateSessionRequest, Manager, NullSessionLocation, Session (+3 more)

### Community 15 - "OIDC Client"
Cohesion: 0.20
Nodes (13): AuthURLOptions, CallbackResult, Client, Config, generateState(), generateUserID(), NewClient(), Provider (+5 more)

### Community 16 - "TOTP Manager"
Cohesion: 0.20
Nodes (8): Context, CredentialStore, Config, Manager, Secret, generateBackupCode(), NewManager(), TestGenerateBackupCode()

### Community 17 - "Noop Auditor"
Cohesion: 0.20
Nodes (12): DefaultAuditor(), NewNoOpAuditor(), BenchmarkNoOpAuditor_Log(), TestDefaultAuditor(), TestNoOpAuditor_Log(), TestNoOpAuditor_LogNil(), NoOpAuditor, AuditEvent (+4 more)

### Community 18 - "Basic Auth Middleware"
Cohesion: 0.23
Nodes (10): GetUser(), NewBasicAuthMiddleware(), BasicAuthConfig, BasicAuthMiddleware, Authenticator, ErrorHandler, Handler, Request (+2 more)

### Community 19 - "JWT Middleware"
Cohesion: 0.24
Nodes (10): GetClaims(), NewJWTMiddleware(), JWTConfig, JWTMiddleware, Claims, ErrorHandler, Handler, Request (+2 more)

### Community 20 - "Session Middleware"
Cohesion: 0.24
Nodes (10): GetSessionData(), NewSessionMiddleware(), SessionConfig, SessionMiddleware, ErrorHandler, Handler, Manager, Request (+2 more)

### Community 21 - "Storage Interfaces"
Cohesion: 0.24
Nodes (10): Time, CredentialStore, OIDCState, OIDCStateStore, SessionData, SessionStore, TokenStore, User (+2 more)

### Community 22 - "Graphify Project Guide"
Cohesion: 0.29
Nodes (7): Community Structure, Cross-File Relationships, God Nodes, graphify, graphify-out, graphify update, graphify-out/wiki/index.md

### Community 23 - "Graphify Query Tools"
Cohesion: 0.40
Nodes (5): graphify-out/graph.json, graphify-out/GRAPH_REPORT.md, graphify explain, graphify path, graphify query

## Knowledge Gaps
- **89 isolated node(s):** `$schema`, `plugin`, `@opencode-ai/plugin`, `Handler`, `AuditLogger` (+84 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **5 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `New()` connect `API Router Auth` to `Basic Auth Tests`, `Memory State Storage`, `Middleware Utilities`, `WebAuthn Credentials`, `OIDC Providers`, `Error Handling`, `Standard Logging`, `Token Tests`, `Token Manager`, `Session Manager`, `OIDC Client`, `TOTP Manager`?**
  _High betweenness centrality (0.460) - this node is a cross-community bridge._
- **Why does `NewTokenManager()` connect `Token Tests` to `Basic Auth Tests`, `API Router Auth`, `Token Manager`?**
  _High betweenness centrality (0.166) - this node is a cross-community bridge._
- **Why does `NewInMemoryUserStore()` connect `Authenticator Tests` to `Basic Auth Tests`, `Memory State Storage`, `Token Tests`, `Session Tests`?**
  _High betweenness centrality (0.092) - this node is a cross-community bridge._
- **Are the 50 inferred relationships involving `NewInMemoryUserStore()` (e.g. with `TestBasicAuthWrapper_Authenticate_Failure()` and `TestBasicAuthWrapper_Authenticate_Success()`) actually correct?**
  _`NewInMemoryUserStore()` has 50 INFERRED edges - model-reasoned connections that need verification._
- **Are the 36 inferred relationships involving `NewInMemoryCredentialStore()` (e.g. with `TestBasicAuthWrapper_Authenticate_Failure()` and `TestBasicAuthWrapper_Authenticate_Success()`) actually correct?**
  _`NewInMemoryCredentialStore()` has 36 INFERRED edges - model-reasoned connections that need verification._
- **Are the 23 inferred relationships involving `New()` (e.g. with `main()` and `NewRouter()`) actually correct?**
  _`New()` has 23 INFERRED edges - model-reasoned connections that need verification._
- **What connects `$schema`, `plugin`, `@opencode-ai/plugin` to the rest of the system?**
  _89 weakly-connected nodes found - possible documentation gaps or missing edges._