### 27. What we'll cover in this section

***

### 28. Getting started with the Logger service
```go
$ go get go.mongodb.org/mongo-driver/mongo
$ go get go.mongodb.org/mongo-driver/options
```

***

### 29. Setting up the Logger data models

***

### 30. Finishing up the Logger data models
```go
$ go get github.com/go-chi/chi/middleware
$ go get github.com/go-chi/cors
```

***

### 31. Setting up routes, handlers, helpers, and a web server in our Logger service

***

### 32. Adding MongoDB to our `docker-compose.yml` file
```go
$ cd logger-service
$ go run ./run/api
```

***

### 33. Add the logger-service to `docker-compose.yml` and the Makefile
```go
$ cd logger-service
$ cd ../project
make up_build
```

***

### 34. Adding a route and handler on the Broker to communicate with the Logger service

***

### 35. Update the front-end to post to the logger, via the Broker

***

### 36. Add basic logging to the Authentication service

***

### 37. Trying things out
* https://www.mongodb.com/products/tools/compass
* __New Connection__:
  * `mongodb:<username>:<password>@localhost:27017/<database>?`
  * `mongodb:admin:password@localhost:27017/logs?authSource=admin&readPreference=primary&appName=MongoDB%20Compass&directConnection=true&ssl=false`

***
