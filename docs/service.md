# Service Module Context

Purpose: legacy service layer. After DDD refactor, prefer `application/*` packages.

Key files:
- `service/user_service.go` - previous user service implementation

Notes:
- Keep for backward compatibility temporarily; migrate callers to `application`.
