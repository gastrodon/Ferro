FROM golang:alpine as builder

WORKDIR /build

COPY . .
RUN go get -u ./...
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o build .

FROM alpine:latest

WORKDIR /build
COPY --from=builder /build/build .

ENV FERROTHORN_CONNECTION ""
ENV FERROTHORN_ROOT ""
ENV FERROTHORN_SECRET ""
ENTRYPOINT ./build
