## 09. What we'll cover in this section
* User (via browser) -> Broker Server (in Docker Image)
  
```
$ Terminal -> New Terminal

$ cd /Users/marshad/Desktop/Golang_Microservices/2_Frontend_and_One_Microservice/go-micro/project
$ make up_build

$ make start
- Browser: localhost
$ make stop
```

***

## 10. Setting up the front-end
```
$ cd /Users/marshad/Desktop/Golang_Microservices/2_Frontend_and_One_Microservice/go-micro/frontend
$ go run ./cmd/web

- Stop: control C
```

***

## 11. Reviewing the front-end code

***

## 12. Our first service: the Broker
```go
$ cd broker-service
$ go mod init broker
$ mkdir -p cmd/api
$ cd cmd/api
$ touch main.go
$ touch routes.go
$ touch handler.go
```

```go
$ go get github.com/go-chi/chi/v5/middleware
$ go get github.com/go-chi/chi/cors
```

```go
$ cd broker-service
$ go run ./cmd/api
```

***

## 13. Building a docker-image for the Broker service
```dockerfile
# base go image
FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

# go build [-o output] [build flags] [packages]
RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

RUN chmod +x /app/brokerApp

###########################
# build a tiny docker-image
###########################

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/brokerApp /app

CMD [ "/app/brokerApp" ]
```

```dockerfile
version: '3'

services:

  broker-service:
    build:
      context: ./../broker-service
      dockerfile: ./../broker-service/broker-service.dockerfile
    restart: always
    ports:
      - "8080:80"
    deploy:
      mode: replicated
      replicas: 1
```

***

## 14. Adding a button and JavaScript to the front-end
```go
$ cd front-end
$ go run ./cmd/web
```

***

## 15. Creating some helper functions to deal with JSON and such

***

## 16. Simplifying things with a Makefile (Mac & Linux)
```dockerfile
$ docker-compose down
$ make stop
$ make up_build
```

***

## 17. Simplifying things with a Makefile (Windows)

***

![Code UML](https://github.com/muarshad01/Microservices-in-Go/blob/main/images/input-output.png)

***

![Code UML](https://github.com/muarshad01/Microservices-in-Go/blob/main/images/broker-code-uml.png)

***

![Code Tree Structure](https://github.com/muarshad01/Microservices-in-Go/blob/main/images/code_tree.png)
***
