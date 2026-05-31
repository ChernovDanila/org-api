FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
COPY vendor/ vendor/
COPY . .
RUN go build -mod=vendor -o main ./cmd/api

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]