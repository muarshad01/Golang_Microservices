### 103 Installing minikube

[minikube start](https://minikube.sigs.k8s.io/docs/start/)

```
$ which minikube
/usr/local/bin/minikube
```
### 104 Installing kubectl

[kubectl](https://kubernetes.io/docs/tasks/tools/)

```
$ which kubectl
/usr/local/bin/kubectl

```

### 105 Initializing a cluster

```
$ minikube start --nodes=2
$ minikube status
$ minikube stop
$ minikube start
$ minikube delete
```

### 106 Bringing up the k8s dashboard

```
$ minikube dashboard
```

### 107 Creating a deployment file for Mongo

```
$ minikube status
$ minikube start
```

```
$ kubectl get pods          # A POD can've one or more containers in it; A POD can've one or more service in it
$ kubectt get pods -A
```

```
$ mkdir k8s
$ cd k8s
$ touch mongo.yml

$ kubectl apply -f k8s
$ kubectl get pods

$ minikube dashboard

$ kubectl get svc

$ kubect get deployments
```

### 108 Creating a deployment file for RabbitMQ

```
$ kubectl apply -f k8s/rabbit.yml
$ kubectl get pods
$ kubectl get svc
```

### 109 Creating a deployment file for Broker service

```
$ kubectl apply -f k8s/broker.yml
```

### 110 When things go wrong

```
$ kubectl get deployments
$ kubectl delete deployments broker-service mongo rabbitmq
$ kubectl get deployments
```

```
$ kubectl get svc
$ kubectl delete svc broker-service mongo rabbitmq
$ kubectl get svc
```

```
$ kubectl apply -f k8s/
$ kubectl get pods
$ kubectl logs broker-service-5cbf7979ff-859sz
```

### 115 Running postgres on the host machine

```
$ docker-compose -f postgres.yml up -d
```

### 117 Trying things out by adding a LoadBalancer service

Three ways: nodePort / LoadBalancer / Ingres (production)

```
$ kubectl get svc
$ kubectl delete svc broker-service
$ kubectl expose deployment broker-service --type=LoadBalancer --port=8080 --target-port=8080
$ kubectl get svc   # EXTERNAL-IP <pending>
$ minikube tunnel
```
