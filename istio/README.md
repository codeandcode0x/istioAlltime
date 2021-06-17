# Istio

## Install Istio

## Install Helm 
```
curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash

brew install helm
```

## Install Nginx-controller
```
kubectl create ns nginx-ingress
helm repo add nginx-stable https://helm.nginx.com/stable
helm repo update
helm install nginx-controller nginx-stable/nginx-ingress -n nginx-ingress

kubectl patch svc ingress-nginx-controller -n nginx-ingress -p '{"spec": {"type": "LoadBalancer", "externalIPs":["node ip"]}}'

```



