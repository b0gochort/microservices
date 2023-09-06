# Dockerfile in build/package directory
FROM golang:1.20 AS builder

WORKDIR /app

# Copy the entire project
COPY / /app

RUN go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o engineservice ./cmd/engineservice/main.go

# Build the final image
FROM alpine:latest

WORKDIR /binary

RUN apk add curl

# Copy the binary from builder
COPY --from=builder /app/engineservice .
COPY --from=builder /app/config.json .


ENTRYPOINT ["./engineservice"]
