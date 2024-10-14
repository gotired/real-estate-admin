FROM golang:1.22.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN mkdir -p /app/temp_sql

COPY ./internal/infrastructure/persistance/database/migrations/*.sql /app/temp_sql/

WORKDIR /app/cmd

RUN go build -o main main.go

FROM ubuntu:latest

COPY --from=builder /app/cmd/main /app/main
COPY --from=builder /app/temp_sql /app/migrations

RUN chmod +x /app/main

ENTRYPOINT ["/app/main"]
