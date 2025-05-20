# Start from the official Golang image for building
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o microservice ./cmd/main.go

# Use a minimal image for running
FROM alpine:latest

# Create a non-root user and group
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# Copy the binary from the builder
COPY --from=builder /app/microservice .

# Change ownership and permissions
RUN chown appuser:appgroup /app/microservice

# Switch to non-root user
USER appuser

# Expose the service port
EXPOSE 8282

# Run the microservice
CMD ["./microservice"]
