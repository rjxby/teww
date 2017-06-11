FROM golang:1.8.1-alpine

RUN mkdir -p $GOPATH/src/github.com/rjxby/teww/teww-backend
ADD . $GOPATH/src/github.com/rjxby/teww/teww-backend
WORKDIR $GOPATH/src/github.com/rjxby/teww/teww-backend

RUN go install .

ENTRYPOINT $GOPATH/bin/teww-backend