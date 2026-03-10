# 59. Kubernetes. NodePort, ClusterIP, LoadBalancer

## 

Use 57/eks-using-modules for EKS cluster.

## Deploy application

```bash
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

## Check service

```bash
kubectl get svc
curl <EXTERNAL-IP>
```