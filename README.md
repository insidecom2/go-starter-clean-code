# go-starter: Clean-ish Go scaffold with Fiber and zerolog

A minimal Go service starter with Postgres, Redis, and Kafka for local development. Provides Docker Compose setup and a production-style multi-stage Dockerfile in /scaffold and at the repo root.

## Quickstart (Docker - recommended)

1. Copy the example env and edit if needed:

   cp .env.docker.example .env

2. Start services (builds the Go image):

   docker compose up --build -d

3. Follow app logs:

   docker compose logs -f app

4. Stop and remove containers:

   docker compose down

## Local quickstart (no Docker)

1. Install dependencies and tidy modules:

   go mod tidy

2. Build and run:

   go build -o bin/app ./cmd/app
   APP_PORT=8080 \
   DATABASE_URL=postgres://postgres:password@localhost:5432/postgres?sslmode=disable \
   REDIS_ADDR=localhost:6379 \
   KAFKA_BROKERS=localhost:9092 \
   ./bin/app

## Services included (docker-compose)

- postgres:15 (db) - port 5432
- redis:7 (cache) - port 6379
- confluentinc cp-zookeeper + cp-kafka (kafka) - ports 9092, 29092
- app (the Go service) - port 8080

## Environment variables

Defaults are provided in `.env.docker.example`:

- APP_PORT (default: 8080)
- POSTGRES_USER / POSTGRES_PASSWORD / POSTGRES_DB
- DATABASE_URL (container default: postgres://postgres:password@db:5432/postgres?sslmode=disable)
- REDIS_ADDR (container default: redis:6379)
- KAFKA_BROKERS (container default: kafka:9092)

Note: use `kafka:9092` for in-container connectivity; from the host use `localhost:29092` (mapped port).

## Endpoints (example)

- GET /api/v1/health
- GET /api/v1/users
- POST /api/v1/users  (body: {"name":"...","email":"..."})

## Development: hot reload with Air

For fast local development use Air (https://github.com/cosmtrek/air) for automatic rebuild & restart on source changes.

Install Air:

- With Go (recommended): go install github.com/cosmtrek/air@latest
- Or on macOS with Homebrew: brew install air

Ensure $GOPATH/bin or $GOBIN is on your PATH, e.g.:

   export PATH=$PATH:$(go env GOPATH)/bin

Prepare environment and services:

1. Copy the example env and adjust for localhost if needed:

   cp .env.docker.example .env

Edit .env to use localhost endpoints if running services locally, e.g.:

   DATABASE_URL=postgres://postgres:password@localhost:5432/postgres?sslmode=disable
   REDIS_ADDR=localhost:6379
   KAFKA_BROKERS=localhost:29092

2. Start required infrastructure (Postgres, Redis, Kafka) with docker compose:

   docker compose up -d db redis zookeeper kafka

Run the app in dev mode (hot reload):

   air

Air will watch source files, rebuild the binary at ./tmp/app and restart it automatically. The included .air.toml builds ./cmd/app by default.

## Tests

Run unit tests:

   go test ./...

## Database access

Open a psql shell to the running Postgres container:

   docker compose exec db psql -U ${POSTGRES_USER:-postgres} -d ${POSTGRES_DB:-postgres}

## Notes and tips

- `scaffold/` contains template Dockerfile and docker-compose.yml kept in sync with the root files.
- Persistent volumes: `db-data`, `redis-data`.
- Add DB migrations as needed; consider adding a one-shot migration service to docker-compose.

---

If you want these changes committed, say so and a commit will be prepared.