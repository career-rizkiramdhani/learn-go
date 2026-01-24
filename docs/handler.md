# Handler Module Context

Purpose: HTTP adapters that translate HTTP requests to application calls and return HTTP responses.

Key files:
- `handler/user.go` - user CRUD endpoints
- `handler/auth.go` - authentication endpoints

Notes:
- Handlers should call functions in `application/*` instead of `service/*` or `models` directly.
