
@test "rboxcli: build rboxcli" {
    cd client
    /usr/local/go/bin/go clean .
    /usr/local/go/bin/go build -o rboxcli
    [ -f rboxcli ]
}

@test "rboxcli: get SSID" {
    cd client
    ./rboxcli get -s $(minikube ip) -p30051 -f ssid 788a20298f81.z3n.com.br
    [ $? -eq 0 ]
}

@test "rboxcli: get MACs" {
    cd client
    ./rboxcli get -s $(minikube ip) -p30051 -f macs 788a20298f81.z3n.com.br
    [ $? -eq 0 ]
}
