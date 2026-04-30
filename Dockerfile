# multistage build
FROM golang:1.25-alpine AS builder
WORKDIR /src

# cache dependencies
COPY go.mod go.sum ./
RUN apk add --no-cache git && go mod download

# build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /app ./cmd/app

FROM gcr.io/distroless/base-debian11
COPY --from=builder /app /app
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app"]
