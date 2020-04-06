FROM golang:1.13

WORKDIR /src
COPY . .

ENV FERRO_SECRET ""
ENV FERRO_LOG_LEVEL 2

RUN mkdir -p /files
RUN go get ./...

ENTRYPOINT go run . -port 80 -at /files -level $FERRO_LOG_LEVEL
