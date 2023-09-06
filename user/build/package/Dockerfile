# Dockerfile in build/package directory
FROM golang:1.20 AS builder

WORKDIR /app

# Copy the entire project
COPY / /app

RUN go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o userservice ./cmd/userservice/main.go

# Build the final image
FROM alpine:latest

WORKDIR /binary

RUN apk add curl

# Copy the binary from builder
COPY --from=builder /app/userservice .
COPY --from=builder /app/config.json .


ENTRYPOINT ["./userservice"]
