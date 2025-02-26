# Build stage
FROM golang:1.24 AS builder
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main_server main.go

# Ensure the binary is executable
RUN chmod +x main_server

# Runtime stage
FROM alpine:latest
WORKDIR /root/

# Install CA certificates for HTTPS/SSL support (needed for RDS)
RUN apk add --no-cache ca-certificates

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main_server .

# Expose the application port
EXPOSE 8080

# Set the correct entrypoint to run the binary with arguments
CMD ["/root/main_server", "start", "--env", "local","--conn","https"]
