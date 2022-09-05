FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app
 
WORKDIR /app

RUN CGO_ENABLED=0 go build -o cashierApp ./cmd/

RUN chmod +x /app/cashierApp

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/cashierApp /app

cmd [ "/app/cashierApp" ] 