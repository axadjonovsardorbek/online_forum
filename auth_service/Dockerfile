# Stage 1: Build the Go binary
FROM golang:1.22.2-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app


# Install git (required for go mod)
RUN apk add --no-cache git

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main ./main.go

# Stage 2: Build a small image
FROM alpine:3.16

WORKDIR /app

RUN mkdir -p /app/logs

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

COPY .env .

# Ensure the binary is executable
RUN chmod +x ./main

# Command to run the executable
CMD ["./main"]
