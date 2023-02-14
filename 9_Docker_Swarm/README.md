
1. Build and tag our documer images and push them somewhere because docker-swarm needs to pull those image from somewhere.

### 76 Building images for our microservices

`https://hub.docker.com/`


```
$ cd logger-service
$ docker build -f logger-service.dockerfile -t marshad1/logger-service:1.0.0 .  #Build & tag
$ docker push marshad1/logger-service:1.0.0


ERROR: denied: requested access to the resource is denied
$ docker login
Username: <username>
Password: <password>

$ cd broker-service
$ docker build -f broker-service.dockerfile -t marshad1/broker-service:1.0.0 .
$ docker push marshad1/broker-service:1.0.0
```
