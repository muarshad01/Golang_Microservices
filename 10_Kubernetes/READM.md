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
