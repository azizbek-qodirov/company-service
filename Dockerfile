# Stage 1: Build the Go binary
FROM golang:1.22.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

COPY config/getEnv.go .

EXPOSE 50051

CMD ["./main"]