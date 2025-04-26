# Stage 1: Build the Go binary
FROM golang:1.23.4 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the Go app (for example, `keystone`)
RUN go build -o keystone .

# Stage 2: Create a minimal, production-ready image
FROM debian:stable-slim

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/keystone .

# Expose port 8080 (or your desired port)
EXPOSE 8080

# Run the Go binary when the container starts
CMD ["./keystone"]