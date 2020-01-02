
FROM golang:1.13

WORKDIR /src
COPY . .

RUN go get ./...

ENTRYPOINT go run . -csr sslkey/server.csr -key sslkey/server.key -at ~/ferro-files/
