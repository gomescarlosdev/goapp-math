FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod init main && \
    go build -o math 

CMD ["./math"]