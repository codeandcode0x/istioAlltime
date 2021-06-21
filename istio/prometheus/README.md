# Prometheus

release download 
```sh
https://prometheus.io/download/
```

install prom
```sh
curl -LO https://github.com/prometheus/prometheus/releases/download/v2.27.1/prometheus-2.27.1.linux-amd64.tar.gz
```

install exporter
```sh
curl -OL https://github.com/prometheus/node_exporter/releases/download/v1.1.2/node_exporter-1.1.2.linux-amd64.tar.gz

tar -xzf node_exporter-1.1.2.linux-amd64.tar.gz

cp node_exporter-1.1.2.linux-amd64/node_exporter /usr/local/bin/

```


Grafana Cloud Stack:

https://grafana.com/orgs/codeandcode0x

Dashboard:
```
Host Stats - Prometheus Node Exporter 0.16.0:  6014
```


kubectl exec -it prometheus-6f558d6655-pmpmh -n istio-system ls /var/run/secrets/kubernetes.io/serviceaccount/