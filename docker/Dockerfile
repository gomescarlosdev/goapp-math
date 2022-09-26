FROM golang:latest

WORKDIR /app

COPY go .

RUN go build -o math

CMD ["./math"]