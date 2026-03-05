# 57. Kubernetes. EKS


```bash
aws eks update-kubeconfig --name miniman-eks-fargate-ua6 --region eu-central-1
# Verify current context
kubectl config current-context
# Check cluster nodes/fargate
kubectl get nodes
# Check all pods across namespaces
kubectl get pods -A
# Check Fargate profiles are ready
aws eks describe-fargate-profile \
  --cluster-name miniman-eks-fargate-ua6 \
  --fargate-profile-name fp-default \
  --region eu-central-1 \
  --query 'fargateProfile.status'
```