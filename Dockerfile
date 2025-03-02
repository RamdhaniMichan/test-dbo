# Build stage
FROM golang:1.20 AS builder
WORKDIR /app

# Copy files and download dependencies
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the binary with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final lightweight image
FROM alpine:latest
WORKDIR /root/

# Install dependencies
RUN apk --no-cache add ca-certificates

# Copy binary
COPY --from=builder /app/main .

# Copy .env file
COPY .env .env

EXPOSE 8080
CMD ["./main"]
