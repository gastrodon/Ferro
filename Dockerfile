
FROM golang:1.13

WORKDIR /src
COPY . .

RUN go get ./...

ENTRYPOINT go run . -crt sslkey/monke-crt.pem -key sslkey/monke-key.pem -at ~/ferro-files/
