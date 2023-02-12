
[gRPC Website](https://grpc.io/)

```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Install VSCode extension "Profobuf support"

## `~/.bash_profile`
```
export GOROOT=/usr/local/go
export GOPATH=$HOME/goworkspace
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH/bin
export GO111MODULE="on"
```

[Protocol Buffer Compiler Installation](https://grpc.io/docs/protoc-installation/)

Goto `https://github.com/protocolbuffers/protobuf/releases` under Assets: download `protoc-21.12-osx-aarch_64.zip`

```
$ ./bin/protoc $GOBIN
```

```
protoc --version
```

