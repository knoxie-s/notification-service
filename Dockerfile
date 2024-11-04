FROM golang:1.22.2-alpine AS builder

COPY . /github.com/knoxie-s/notification-service
WORKDIR /github.com/knoxie-s/notification-service

RUN go mod download
RUN go build -o ./bin/notification-service cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY .env .
COPY --from=builder /github.com/knoxie-s/notification-service/bin/notification-service .

CMD ["./notification-service"]