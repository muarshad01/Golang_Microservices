
1. Build and tag our documer images and push them somewhere because docker-swarm needs to pull those image from somewhere.

### 76 Building/Tagging images for our microservices

`https://hub.docker.com/`


```
$ cd logger-service
$ docker build -f logger-service.dockerfile -t marshad1/logger-service:1.0.0 .  #Build & tag
$ docker push marshad1/logger-service:1.0.0


ERROR: denied: requested access to the resource is denied
$ docker login
Username: <username>
Username: marshad1
Password: <password>

$ cd broker-service
$ docker build -f broker-service.dockerfile -t marshad1/broker-service:1.0.0 .
$ docker push marshad1/broker-service:1.0.0
```

### 78 Initializing and starting Docker Swarm

```
$ docker swarm init
```

```
$ docker swarm join-token <role>
$ docker swarm join-token worker
$ docker swarm join-token manager
```

Deploy Docker Swarm 

```
$ docker stack deploy -c <configuration-file> <swarn-name>
$ docker stack deploy -c swarm.yml myapp

$ docker service ls
```

### 80 Scaling services

```
$ docker service scale myapp_listener-service=3
$ docker service ls
```

### 81 Updating services

```
$ cd logger-service
$ docker build -f logger-service.dockerfile -t marshad1/logger-service:1.0.1 .
$ docker push marshad1/logger-service:1.0.1
$ docker service ls
```

```
$ docker service update --image marshad1/logger-service:1.0.1 myapp_docker-service #up-grade
$ docker service ls
$ docker service update --image marshad1/logger-service:1.0.0 myapp_docker-service #down-grade
$ docker service ls
```

### Stopping docker swarm

```
$ docker stack rm myapp
$ docker swarm leave --force
```
