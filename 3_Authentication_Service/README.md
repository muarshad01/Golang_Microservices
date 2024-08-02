## 18. What we'll cover in this section
* User try to Authenticate through the Broker
* Persistence (Postgres db)
  
***

## 19. Setting up a stub Authentication service
```go
$ cd authentication-service
$ go mod init
```

```go
$ go get github.com/go-chi/chi/v5/middleware
$ go get github.com/go-chi/cors
```

***

## 20. Creating and connecting to Postgres from the Authentication service
```go
$ go get github.com/jackc/pgconn
$ go get github.com/pgx/v4
$ go get github.com/pgx/stdlib
```
***

## 21. A note about PostgreSQL
```dockerfile
postgres:
    image: 'postgres:14.2'
    ports:
      - "5432:5432"
    deploy:
      mode: replicated
      replicas: 1
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: password
      POSTGRES_DB: users
    volumes:
      - ./db-data/postgres/:/var/lib/postgresql/data/
```

```
$ mkdir ./project/db-data/postgres
```
***

## 22. Updating our `docker-compose.yml` for Postgres and the Authentication service

```
FROM alpine:latest

RUN mkdir /app

COPY authApp /app

CMD ["/app/authApp"]
```

```go
$ cd <...>/project
$ make up_build
```
***

## 23. Populating the Postgres database

***

## 24. Adding a route and handler to accept JSON

***

## 25. Update the Broker for a standard JSON format, and conect to our Auth service

***

## 26. Updating the front end to authenticate thorough the Broker and trying things out

***

## JSON Toolbox
* [BeekeeperStudio Client for Postgres](https://www.beekeeperstudio.io/)
* [JSON Toolboc](https://github.com/tsawler/toolbox)

***

## Postgres
```
$ go get github.com/jackc/pgconn
$ go get github.com/jackc/pgx/v4
$ go get github.com/jackc/pgx/v4/stdlib
```

***
