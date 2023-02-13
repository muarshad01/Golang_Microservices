
[gRPC Website](https://grpc.io/)

***

## Tools for compiling <file-name>.proto file

```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27
$ go install google.golang.org/grpc/    cmd/protoc-gen-go-grpc@v1.2
```

***

## Install VSCode extension 

`Profobuf support`

***

## `~/.bash_profile`

```bash
export GOROOT=/usr/local/go
export GOPATH=$HOME/goworkspace
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH/bin
export GO111MODULE="on"
```

***

[Protocol Buffer Compiler Installation](https://grpc.io/docs/protoc-installation/)

Goto --> `https://github.com/protocolbuffers/protobuf/releases` --> Assets `protoc-21.12-osx-aarch_64.zip`

```
$ cp ./bin/protoc $GOBIN
$ protoc --version
```

***

```go
$ protoc \ 
--go_out=. \ 
--go_opt=paths=source_relative \
--go-grpc_out=. \
--go-grpc_opt=paths=source_relative \ 
logs.proto
```

***

```
$ go get google.golang.org/grpc
$ go get google.golang.org/protobuf
```
