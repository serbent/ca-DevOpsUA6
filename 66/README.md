# 66 Helm. Creating a chart


```bash
helm install <RELEASE_NAME> <CHART_NAME>
```

```bash
helm upgrage <RELEASE_NAME> <CHART_NAME>
```

```bash
helm upgrage <RELEASE_NAME> <CHART_NAME> -f values-file.yaml
```

```bash
kubectl port-forward svc/service_name LOCAL_PORT:SERVICE_PORT
```

```bash
kubectl get secrets | grep helm
```

```bash
  kubectl get secret secret_name -o jsonpath='{.data.release}' | \
  base64 --decode | base64 --decode | gunzip
```