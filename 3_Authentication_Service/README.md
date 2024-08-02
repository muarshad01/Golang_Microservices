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
$ go get github.com/go-chi/chi/v5
$ go get github.com/go-chi/chi/v5/middleware
$ go get github.com/go-chi/chi/cors
```

***

## 20. Creating and connecting to Postgres from the Authentication service

***

## 21. A note about PostgreSQL

***

## 22. Updating our docker-compose.yml for Postgres and the Authentication service

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
