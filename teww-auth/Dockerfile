FROM golang:1.8.1-alpine

RUN mkdir -p $GOPATH/src/github.com/rjxby/teww/teww-auth
ADD . $GOPATH/src/github.com/rjxby/teww/teww-auth
WORKDIR $GOPATH/src/github.com/rjxby/teww/teww-auth

RUN go install .

ENTRYPOINT $GOPATH/bin/teww-auth