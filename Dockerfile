
FROM golang:1.13

WORKDIR /src
COPY . .

RUN go build -o cdn .
ENTRYPOINT ./cdn -csr sslkey/server.csr -key sslkey/server.key -at ~/ferro-files/ -port 443
