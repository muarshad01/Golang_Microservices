### 103. Installing `Minikube`
* [minikube start](https://minikube.sigs.k8s.io/docs/start/)
```bash
$ which minikube
/usr/local/bin/minikube
```

***

### 104. Installing `Kubectl`
* [kubectl](https://kubernetes.io/docs/tasks/tools/)
```bash
$ which kubectl
/usr/local/bin/kubectl
```

***

### 105. Initializing a Cluster
```bash
$ minikube start --nodes=2
$ minikube status
$ minikube stop
$ minikube start
$ minikube delete
```

***

### 106. Bringing up the k8s Dashboard
```bash
$ minikube dashboard
```

***

### 107. Creating a Deployment file for MongoDB
```bash
$ minikube status
$ minikube start
```

```bash
$ kubectl get pods          # A POD can've one or more containers in it; A POD can not have one or more service in it
$ kubectt get pods -A
```

```bash
$ mkdir k8s
$ cd k8s
$ touch mongo.yml

$ kubectl apply -f k8s
$ kubectl get pods

$ minikube dashboard

$ kubectl get svc

$ kubect get deployments
```

***

### 108. Creating a Deployment file for RabbitMQ
```bash
$ kubectl apply -f k8s/rabbit.yml
$ kubectl get pods
$ kubectl get svc
```

***

### 109 Creating a Deployment file for broker-service
```bash
$ kubectl apply -f k8s/broker.yml
```
***

### 110 When things Go Wrong
```bash
$ kubectl get    deployments
$ kubectl delete deployments broker-service mongo rabbitmq
$ kubectl get    deployments
```

```bash
$ kubectl get    svc
$ kubectl delete svc broker-service mongo rabbitmq
$ kubectl get    svc
```

```bash
$ kubectl apply -f k8s/
$ kubectl get pods
$ kubectl logs broker-service-5cbf7979ff-859sz
```

***

### 115 Running EXTERNAL Postgres on the `host` machine
```bash
$ docker-compose -f postgres.yml up -d
```

***

### 117 Trying things out by adding a `LoadBalancer` Service
```bash
$ kubectl get svc
```

* None of these ports are available to outside world. Everything is inside the cluster. Expose at-least one service from that cluster to the local-network.

* Three ways
  * `nodePort`
  * `LoadBalancer`
  * `Ingres (production)`

```
$ kubectl get    svc
$ kubectl delete svc broker-service
$ kubectl expose deployment broker-service --type=LoadBalancer --port=8080 --target-port=8080
$ kubectl get svc      # EXTERNAL-IP <pending>
$ minikube tunnel
```

***

### 118 Creating a Deployment file for Front-end Microservice  
```bash
$ kubectl apply -f k8s/front-end.yml
```

***

### 119 Adding an nginx Ingress to our Cluster
```bash
$ minikube addons enable ingress
```

***

### 120 Trying out our Ingress
```bash
$ kubectl apply -f ingress.yml 
$ kubect get ingress
```
* [How To Configure Ingress TLS/SSL Certificates in Kubernetes](https://devopscube.com/configure-ingress-tls-kubernetes/)

***
