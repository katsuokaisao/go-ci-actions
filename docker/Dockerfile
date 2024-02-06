FROM golang:1.20.0-bullseye as builder

WORKDIR /app

COPY . .

RUN go build -o main main.go

FROM debian:bullseye-slim

COPY --from=builder /app/main /app/main

CMD ["/app/main"]
