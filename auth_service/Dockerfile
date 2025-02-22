FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o auth-service main.go
FROM alpine:latest
WORKDIR /root/

# Copy the compiled binary
COPY --from=builder /app/auth-service .

# Expose the port
EXPOSE 8080

# Start the service
CMD ["./auth-service", "--env", "local", "--conn", "https"]