# 37. Logging


Create a kubernetes namespace: 
```bash
kubectl ns create observability
```

Installation of alloy in kubernetes:

```bash
helm upgrade -i alloy grafana/alloy -n observability -f alloy.yml
```

Installation of loki in kubernetes:
```bash
helm upgrade -i loki grafana/loki -n observability -f loki.yaml
```

