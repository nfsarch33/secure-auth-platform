# Build Stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy module files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server/main.go

# Production Stage
FROM alpine:3.18

WORKDIR /root/

# Install runtime dependencies (e.g. ca-certificates for HTTPS)
RUN apk add --no-cache ca-certificates

# Copy binary from builder
COPY --from=builder /app/main .

# Expose port
EXPOSE 8080

# Command to run
CMD ["./main"]
