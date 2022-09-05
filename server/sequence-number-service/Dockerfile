FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app
 
WORKDIR /app

RUN CGO_ENABLED=0 go build -o sequenceNumberApp ./cmd/

RUN chmod +x /app/sequenceNumberApp

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/sequenceNumberApp /app

cmd [ "/app/sequenceNumberApp" ] 