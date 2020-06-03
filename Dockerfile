FROM golang:alpine 

ADD ./src /go/src/podSetLogger
WORKDIR /go/src/podSetLogger

COPY go.mod .
COPY go.sum .

ENV PORT=3001

CMD go run main.go