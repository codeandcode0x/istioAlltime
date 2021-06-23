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
Kubernetes cluster monitoring : 315
```


kubectl exec -it prometheus-6f558d6655-pmpmh -n istio-system ls /var/run/secrets/kubernetes.io/serviceaccount/



container_cpu_usage_seconds_total{image!="",pod!=""}

过滤1分钟内的数据:
rate(container_cpu_usage_seconds_total{image!="",pod!=""}[1m])

使用sum函数,pod在1分钟内的使用率，同时将pod名称打印出来:
sum by (pod)(rate(container_cpu_usage_seconds_total{image!="", pod!=""}[1m] ))


apiserver在1分钟内请求的数
sum(rate(apiserver_request_duration_seconds_sum[1m]))


sum(kube_pod_container_resource_requests_memory_bytes) by (namespace, pod, node)
  * on (pod) group_left()  (sum(kube_pod_status_phase{phase="Running"}) by (pod, namespace) == 1)