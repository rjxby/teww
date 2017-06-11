FROM golang:1.8.1-alpine

RUN mkdir -p $GOPATH/src/github.com/rjxby/teww/teww-client
ADD . $GOPATH/src/github.com/rjxby/teww/teww-client
WORKDIR $GOPATH/src/github.com/rjxby/teww/teww-client

RUN go install .

ENTRYPOINT $GOPATH/bin/teww-client