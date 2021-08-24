FROM golang:1.17.0-bullseye

EXPOSE 50051

RUN adduser rbox

WORKDIR /go/src/github.com/ronaldoafonso/rbox

USER rbox:rbox

COPY --chown=rbox:rbox . .

RUN go get -d -v ./...

RUN cd server && \
    go build -o $GOPATH/bin/rbox

CMD ["rbox"]
