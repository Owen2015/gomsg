# set GOPATH /home/owen/workspace/go
mkdir -p $GOPATH/bin
mkdir -p $GOPATH/pkg
mkdir -p $GOPATH/src

go get github.com/gorilla/websocket

mkdir -p $GOPATH/src/golang.org/x
cd $GOPATH/src/golang.org/x

# git clone https://github.com/golang/net.git
