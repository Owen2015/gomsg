# gomsg

## purpose
this project is a test project to build a message middleware for real time stock data publishing.

## function
it will receive data sent from some stock data generator and push data to clients who have scribe it.

## deploy
this project is writen in go, so  you may need to set up the go environment first.
###  linux deploy
#### Go download and install
1. you need to download Go and install it, using "wget https://dl.google.com/go/go1.10.linux-amd64.tar.gz" to download Go.
2. install Go using "tar -C /usr/local -xzf go1.10.linux-amd64.tar.gz"
3. export path environment variable using "export PATH=$PATH:/usr/local/go/bin"
if something wrong happen, refer to Go's official website for more information. 
https://golang.org/doc/install

#### Go workspace setup
after you install Go environment, you need to set up a go workspace.
1. find a suitable place for your go workspace and create a go workspace hierarchy using "mkdir -p go/bin","mkdir -p go/pkg", "mkdir -p go/src"
2. set a go workspace environment varaible "GOPATH" using "set GOPATH YourGoWorkspace" and export environment varaible using "export PATH=$PATH:$GOPATH/bin"


if something wrong happen, refer to Go's offical website for more information.
https://golang.org/doc/code.html

#### third part libarary install
Since the project use a third part liberary gorrila, we need to install it.
enter you go workspace and using command "go get github.com/gorilla/websocket" to install it.

#### project install and running
1. Enter your source directory of go workspace and download the project using "git@github.com:Owen2015/gomsg.git"
2. You can change project's server port or request path by delve into main.go source file and change it. Since only a few lines in there, it will not be hard.
3. Enter go workspace and using "go install yourpackagepath/gomsg" to install it.
4. Now there should be a gomsg executable file in you bin folder. Using "./bin/gomsg" to run it
5. Open browser and enter "http://localhost:3000/" and you can access the stock data in real time.

