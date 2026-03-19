https://oidc.eks.eu-central-1.amazonaws.com/id/E01C9B728F4A7B1EF60185B84F9CCE9D

```
oidc_id=$(aws eks describe-cluster --name ca-devops-ua6 --query "cluster.identity.oidc.issuer" --output text | cut -d '/' -f 5)
```

```
aws iam list-open-id-connect-providers | grep $oidc_id | cut -d "/" -f4
```

```
aws iam create-role \
  --role-name AmazonEKSLoadBalancerControllerRole \
  --assume-role-policy-document file://"load-balancer-role-trust-policy.json"
```

```
aws iam attach-role-policy \
  --policy-arn arn:aws:iam::034362059178:policy/AWSLoadBalancerControllerIAMPolicy \
  --role-name AmazonEKSLoadBalancerControllerRole
```


```
aws eks update-kubeconfig --region eu-central-1 --name ca-devops-ua6
```

```
kubectl apply \
    --validate=false \
    -f https://github.com/jetstack/cert-manager/releases/download/v1.13.5/cert-manager.yaml
```

```
curl -Lo v2_7_1_full.yaml https://github.com/kubernetes-sigs/aws-load-balancer-controller/releases/download/v2.7.1/v2_7_1_full.yaml
```


```
sed -i.bak -e '596,604d' ./v2_7_1_full.yaml
```

```
sed -i.bak -e 's|your-cluster-name|ca-devops-ua6|' ./v2_7_1_full.yaml
```

```
kubectl apply -f v2_7_1_full.yaml
```

```
curl -Lo v2_7_1_ingclass.yaml https://github.com/kubernetes-sigs/aws-load-balancer-controller/releases/download/v2.7.1/v2_7_1_ingclass.yaml
```

```
kubectl apply -f v2_7_1_ingclass.yaml
```

```
kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/v2.7.2/docs/examples/2048/2048_full.yaml
```

```
kubectl get ingress/ingress-2048 -n game-2048
```

