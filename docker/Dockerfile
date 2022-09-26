FROM golang:1.15.15

WORKDIR /app

COPY go .

RUN go build -o math

CMD ["./math"]