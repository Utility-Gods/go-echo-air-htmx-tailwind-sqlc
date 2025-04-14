# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git make

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Install tools and build the application
RUN go install github.com/a-h/templ/cmd/templ@latest && \
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest && \
    go generate ./... && \
    CGO_ENABLED=0 GOOS=linux go build -o /app/bin/photogo ./cmd/main.go

# Download Tailwind CSS CLI
RUN wget -O /app/bin/tailwindcss https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.1/tailwindcss-linux-x64 && \
    chmod +x /app/bin/tailwindcss

# Build CSS
RUN /app/bin/tailwindcss -i ./static/css/input.css -o ./static/css/output.css --minify

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy binary and static files
COPY --from=builder /app/bin/photogo /app/bin/photogo
COPY --from=builder /app/static /app/static
COPY --from=builder /app/internal/database/migrations /app/internal/database/migrations

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Set environment variables
ENV TZ=UTC \
    PORT=6969

# Run the application
CMD ["/app/bin/photogo"] 