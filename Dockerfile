FROM golang:1.13

WORKDIR /src
COPY . .

ENV MONGO_HOST "localhost"
ENV MONGO_USER "admin"
ENV MONGO_PASS ""

ENV FERRO_SECRET ""

RUN mkdir -p /files
RUN go get ./...
RUN go test ./...

ENTRYPOINT go run . -port 80 -at /files
