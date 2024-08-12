### 66. What we'll cover in this section

***

### 67. Installing the necessary tools for `gRPC`
```go
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

***

### 68. Defining a Protocol for gRPC: the `.proto` file
```go
$ cd logger-service
$ mkdir logs & cd logs
$ touch logs.proto
```

```proto
syntax = "proto3";

package logs;

option go_package = "/logs";

message Log {
    string name = 1;
    string data = 2;
}

message LogRequest {
    Log logEntry = 1;
}

message LogResponse {
    string result = 1;
}

service LogService {
    rpc WriteLog(LogRequest) returns (LogResponse);
}
```

***

### 69. Generating the gRPC code from the command line
```go
$ cd ~/Desktop/Golang_Microservices/8_gRPC/logger-service/logs
```

#### [Protocol Buffer Compiler Installation](https://grpc.io/docs/protoc-installation/)
* https://github.com/protocolbuffers/protobuf/releases
* Assets `protoc-21.12-osx-aarch_64.zip`

```bash
$ cp ./bin/protoc $GOBIN
$ protoc --version
```

```go
$ protoc \ 
--go_out=. \ 
--go_opt=paths=source_relative \
--go-grpc_out=. \
--go-grpc_opt=paths=source_relative \ 
logs.proto
```

```bash
- logs.pb.go
- logs.proto
```

***

### 70. Getting started with the gRPC server
```go
$ go get google.golang.org/grpc
$ go get google.golang.org/protobuf
```

```go
package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"log-service/logs"
	"net"

	"google.golang.org/grpc"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	// write the log
	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{Result: "failed"}
		return res, err
	}

	// return response
	res := &logs.LogResponse{Result: "logged!"}
	return res, nil
}

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	logs.RegisterLogServiceServer(s, &LogServer{Models: app.Models})

	log.Printf("gRPC Server started on port %s", gRpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
```

***

### 71. Listening for gRPC connections in the Logger microservice

***

### 72. Writing the client code

***

### 73. Updating the front end code

***

### 74. Trying things out

***

* [gRPC Website](https://grpc.io/)

***

### Install VSCode extension 
* Profobuf support

***

### `~/.bash_profile`
```bash
export GOROOT=/usr/local/go
export GOPATH=$HOME/goworkspace
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH/bin
export GO111MODULE="on"
```
***
