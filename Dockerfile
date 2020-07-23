FROM golang:latest

WORKDIR /src
COPY . .

ENV FERROTHORN_CONNECTION ""
ENV FERROTHORN_ROOT ""
ENV FERROTHORN_SECRET ""

RUN go get ./...
RUN go build -o ferrothorn .

ENTRYPOINT ./ferrothorn
