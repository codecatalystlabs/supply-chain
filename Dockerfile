### Multi-stage Dockerfile for MoH supply-chain backend

## 1) Builder image: compile Go binary
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install git (needed by go modules that use VCS)
RUN apk add --no-cache git

# Pre-copy go.mod/go.sum to leverage Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the server binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server


## 2) Runtime image: minimal Alpine with compiled binary
FROM alpine:3.19

WORKDIR /app

# Create non-root user
RUN adduser -D -g '' appuser

# Copy compiled binary and required assets
COPY --from=builder /app/server ./server
COPY --from=builder /app/internals/web ./internals/web
COPY --from=builder /app/docs ./docs

# Environment (can be overridden at runtime)
ENV APP_PORT=5500 \
    GIN_MODE=release

EXPOSE 5500

USER appuser

# Start the HTTP server
CMD ["./server"]

