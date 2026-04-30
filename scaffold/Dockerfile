# multistage build
FROM golang:1.21-alpine AS builder
WORKDIR /src
COPY . .
RUN apk add --no-cache git \
    && go build -o /app ./cmd/app

FROM gcr.io/distroless/base-debian11
COPY --from=builder /app /app
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app"]
