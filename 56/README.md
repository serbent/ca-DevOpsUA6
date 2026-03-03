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
kubectl get nodes -o wide 
kubectl get pods
kubectl get services
kubectl get deployments
kubectl get replicasets
kubectl get pods
```