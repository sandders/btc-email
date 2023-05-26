FROM golang:1.20-bullseye

WORKDIR /app

COPY . .

RUN go build -o btc-email

EXPOSE 8080

CMD ["./btc-email"]
