# 60. ReplicaSet, ReplicationController, DaemonSet


To install metrics server
```bash
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
```

Curl infinite onliner:
```bash
URL=http://aa922273b6e4c4039999b19d18ff8e3b-1581112053.eu-central-1.elb.amazonaws.com
while true; do echo -n "$(date +%H:%M:%S) "; curl -s -o /dev/null -w "%{http_code} %{time_total}s\n" $URL; done
```