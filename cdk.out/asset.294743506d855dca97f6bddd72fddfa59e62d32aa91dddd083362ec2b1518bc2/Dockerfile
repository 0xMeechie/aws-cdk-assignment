FROM golang:1.20.2-alpine3.16 AS builder
WORKDIR /app
COPY . . 
RUN go build worker-main.go

## 
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/worker-main /app

EXPOSE 8080
CMD ["/app/worker-main"]
