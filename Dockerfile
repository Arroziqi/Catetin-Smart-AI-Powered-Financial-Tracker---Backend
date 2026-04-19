# Build stage
FROM golang:1.25.6-alpine AS builder

WORKDIR /app

# Install git
RUN apk add --no-cache git

# Copy go mod
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build binary
RUN go build -o main ./cmd/main.go

# Run stage
FROM alpine:latest

WORKDIR /app

# Copy binary dari builder
COPY --from=builder /app/main .

# Copy .env (optional)
COPY .env .

EXPOSE 3000

CMD ["./main"]
