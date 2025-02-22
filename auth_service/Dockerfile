# Build stage
FROM golang:1.24 AS builder
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build the Go binary
RUN go build -o auth-service main.go

# Runtime stage
FROM alpine:latest
WORKDIR /root/

# Install CA certificates for HTTPS/SSL support (needed for RDS)
RUN apk add --no-cache ca-certificates

# Copy the compiled binary from the builder stage
COPY --from=builder /app/auth-service .

# Expose the application port
EXPOSE 8080

# Set the correct entrypoint to run the binary with arguments
CMD ["/root/auth-service", "--env", "dev", "--conn", "https"]
