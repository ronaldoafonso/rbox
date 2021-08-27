
FROM golang:1.17.0-bullseye AS devel

RUN apt update && \
    apt install -y bats

RUN adduser rbox

WORKDIR /go/src/github.com/ronaldoafonso/rbox

USER rbox:rbox

COPY --chown=rbox:rbox . .

RUN go get -d -v ./...

RUN cd $GOPATH/src/github.com/ronaldoafonso/rbox/server && \
    go build -o $GOPATH/bin/rbox && \
    cd $GOPATH/src/github.com/ronaldoafonso/rbox/client && \
    go build -o $GOPATH/bin/rboxcli



FROM golang:1.17.0-bullseye

EXPOSE 50051

RUN adduser rbox

WORKDIR /home/rbox

USER rbox:rbox

COPY --from=devel $GOPATH/bin/rbox $GOPATH/bin

COPY --from=devel $GOPATH/src/github.com/ronaldoafonso/rbox/.ssh /home/rbox/.ssh

COPY --from=devel $GOPATH/src/github.com/ronaldoafonso/rbox/rbox/uci /home/rbox/uci

CMD ["rbox"]
