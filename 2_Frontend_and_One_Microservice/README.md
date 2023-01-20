## Project

* Broker service 
  - in a docker image.
* Any changes made to Broker code *automatically complied into a new Docker image* and started in Docker.
* Frontend sends a request to broker, which processes the request and sends the response back.

***

```
$ mkdir go-micro
$ mv frontend go-micro
```

* File -> Save Workspace AS
  - Save As: **workspace.code-workspace**
  - <...>/**go-micro**
* File -> Add Folder to Workspace
  - <...>/**go-micro/front-end**

```
$ go run ./cmd/web
```

## Broker

```
$ go get github.com/go-chi/chi/v5
$ go get github.com/go-chi/chi/v5/middleware
$ go get github.com/go-chi/cors
```
