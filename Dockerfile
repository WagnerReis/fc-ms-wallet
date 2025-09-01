FROM golang:1.24-alpine

WORKDIR /app/

RUN apk add --no-cache librdkafka-dev

CMD ["tail", "-f", "/dev/null"]