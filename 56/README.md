# 56. Kubernetes

## Minikube

```bash
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
mv minikube-linux-amd64 minikube
chmod +x minikube
sudo mv minikube /usr/local/bin/
```

## kind

```yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  image: kindest/node:v1.35.0
- role: worker
  image: kindest/node:v1.35.0
- role: worker
  image: kindest/node:v1.35.0
```
```bash
kind create cluster --config kind-config.yaml
```

##  Basic Kubernetes Commands

```bash
alias k='minikube kubectl'
k get nodes -o wide 
k get pods
k get services
k get deployments
k get replicasets
k get pods
k get all -A # all resources in all namespaces
k get pods -o wide # all resources in all namespaces with wide output
k apply -f deployment.yaml
k apply -f service.yaml
k exec my-app-7dfddd84d9-6vbsv bash
k port-forward svc/my-app 8080:80
```