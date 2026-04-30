Go Fiber + Postgres starter (clean-ish architecture) in ./scaffold

Quick start:
1. cp .env.example .env and update DATABASE_URL
2. cd scaffold
3. go mod tidy
4. go run ./cmd/app

Next: add migrations (golang-migrate), repository layer, services, JWT auth, Dockerfile, CI and metrics.
