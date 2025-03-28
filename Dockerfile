# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install build essentials
RUN apk add --no-cache git make build-base

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies with verbose output
RUN go mod download && go mod verify

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .

# Set the binary as the entrypoint
ENTRYPOINT ["./main"]