go-starter: Clean-ish Go scaffold with Fiber and zerolog

Quickstart:

1. Install deps and tidy: go mod tidy
2. Run server: go run ./cmd/server

Endpoints:
- GET /api/v1/health
- GET /api/v1/users
- POST /api/v1/users  (body: {"name":"...","email":"..."})

Notes:
- Uses a threadsafe in-memory repo for demo. Replace with SQL repository in internal/user/repo.
- Put secrets in environment variables or a secret manager.
