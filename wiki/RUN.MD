# Run

## go build 

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o
```

## mariadb
docker run -p 127.0.0.1:3306:3306  --name mariadb -e MARIADB_ROOT_PASSWORD=root123 -d mariadb:10.2.38

## backend
docker run -p 8080:8080  --name ticket-manager -d roandocker/ticket-manager:0.0.1

## run

docker-compose -f docker-compose.yml -p ticket up

## go module off
GO111MODULE=off go get -v github.com/swaggo/swag/cmd/swag

minikube
```
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube

minikube start --driver=none

minikube config set cpus 12
minikube config set memory 28000

```

## Jaeger
```sh
docker pull ethansmart-docker.pkg.coding.net/istioalltime/roandocker/jaegertracing-all-in-one:1.22.0


docker run \
-p 5775:5775/udp \
-p 16686:16686 \
-p 6831:6831/udp \
-p 6832:6832/udp \
-p 5778:5778 \
-p 14268:14268 \
--name=jaegertracing -it ethansmart-docker.pkg.coding.net/istioalltime/roandocker/jaegertracing-all-in-one:1.22.0



```
