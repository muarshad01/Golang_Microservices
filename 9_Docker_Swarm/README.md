
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

### 82 Stopping docker swarm

```
$ docker stack rm myapp
$ docker swarm leave --force
```

### 86 Adding Caddy to the mix as a Proxy to our front end and the broker

Apache / Nginx / Caddy

[Caddy](https://caddyserver.com/)

```
$ docker build -f caddy.dockerfile -t marshad1/micro-caddy:1.0.0 .
$ docker push marshad1/micro-caddy:1.0.0
```

### 87 Modifying out hosts file...

```
$ docker swarm init
$ docker stack deploy -c swarm.yml myapp
```

### 88 Correcting the URL

### 89 Solution to challenge

```
$ make build_front_linux
$ cd ../front-end
$ docker build -f front-end.dockerfile -t marshad1/front-end:1.0.1 .
$ docker push marshad1/front-end:1.0.1
```

### 92 Setting up a non-root account and putting a firewall in place


```
# Create a user
$ ssh root@a.b.c.d
$ adduser mua
New password:
Retype new password:

# Sudo priviliges
$ usermod -aG sudo mua

# Setup basic firewall
$ ufw allow ssh     # uncomplicated fire-wall
$ ufw allow http
$ ufw allow https
$ ufw allow 2377/tcp
$ ufw allow 7946/tcp
$ ufw allow 7946/udp
$ ufw allow 4789/udp
$ ufw allow 8025/tcp

$ ufw enable
$ ufw status
```

### 93 Installing Docker on the servers
