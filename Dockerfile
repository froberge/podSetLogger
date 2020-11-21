FROM golang:alpine

ARG confFilePah

ADD ./src /go/src/podSetLogger

ADD ./config $confFilePah
WORKDIR /go/src/podSetLogger

COPY go.mod .
COPY go.sum .

ENV PORT=3001

CMD go run main.go
