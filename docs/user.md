# User Module Context

Purpose: user domain, repository, application services and HTTP handlers for CRUD operations.

Overview:
- Domain: `domain/user` defines the `User` entity and the `Repository` interface.
- Application: `application/user` contains use-case implementations (Create, Read, Update, Delete, List).
- Infrastructure: `infrastructure/persistence/gorm_user_repository.go` implements the `Repository` interface using GORM models (`models/user.go`).
- Handler: `handler/user.go` exposes HTTP endpoints and maps request/response DTOs in `dto/user.go`.

Endpoints (registered in `router/root.go` under `/api`):
- POST `/api/users` — create user (body: `CreateUserRequest`)
- GET `/api/users` — list users (query: `page`, `size`)
- GET `/api/users/:id` — get user by id
- PUT `/api/users/:id` — update user
- DELETE `/api/users/:id` — delete user

Domain responsibilities:
- Keep `domain/user` free from infrastructure and application concerns.
- Define the `Repository` interface used by application services.

Application responsibilities:
- Validate uniqueness and business invariants.
- Orchestrate hashing password and setting audit fields.
- Call the domain repository to persist data.

Infrastructure responsibilities:
- Implement data mapping between ORM models (`models`) and domain entities.
- Keep GORM-specific tags inside `models/` only.

Notes on CRUD behavior:
- Create: hashes password and generates a user `Code` (timestamp-based by default).
- Read: returns domain `User` mapped to DTOs; 404 if not found.
- Update: optional fields allowed; password update re-hashes password.
- Delete: performs soft-delete via GORM `DeletedAt` field (if models support it).

Testing:
- To test the user module, run the server and exercise endpoints with `curl` or Postman.
- Consider adding unit tests for `application/user` using a mock `domain/user.Repository`.

Examples

Create user (curl):

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"username":"alice","password":"secret123","email":"alice@example.com"}'
```

List users:

```bash
curl http://localhost:8080/api/users?page=1&size=10
```
