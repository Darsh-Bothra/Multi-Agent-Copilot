FROM golang:1.24-alpine AS base
WORKDIR /app
COPY apps/api/go.mod apps/api/go.sum ./
RUN go mod download

FROM base AS dev
# Copy the API source for local development workflows.
COPY apps/api/ ./
EXPOSE 8080
CMD ["go", "run", "./cmd/server"]

FROM base AS builder
# Build a small Linux binary for production runtime image.
COPY apps/api/ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/server ./cmd/server

FROM alpine:3.21 AS prod
WORKDIR /app
RUN adduser -D -g '' appuser
COPY --from=builder /out/server /app/server
EXPOSE 8080
USER appuser
CMD ["/app/server"]

FROM base AS dev
COPY apps/api/ ./
RUN go install github.com/air-verse/air@latest
EXPOSE 8080
CMD ["air", "-c", ".air.toml"]