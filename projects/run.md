# Run

## build 

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o


## mariadb
docker run -p 127.0.0.1:3306:3306  --name mariadb -e MARIADB_ROOT_PASSWORD=root123 -d mariadb:10.2.38

## backend
docker run -p 8080:8080  --name ticket-manager -d roandocker/ticket-manager:0.0.1

## run

docker-compose -f docker-compose.yml -p ticket up

## go module off
GO111MODULE=off go get -v github.com/swaggo/swag/cmd/swag


## grpc
```
protoc -I ${protoDir}/ ${protoDir}/*proto --go_out=plugins=grpc:${outDir}

```

example:

golang
```sh
protoc -I rpc/grpc/protos/movie/ rpc/grpc/protos/movie/*proto --go_out=plugins=grpc:rpc/grpc/protos/movie
```

java
```sh
protoc --plugin=protoc-gen-grpc-java \
  --grpc-java_out="$OUTPUT_FILE" --proto_path="$DIR_OF_PROTO_FILE" "$PROTO_FILE"

mvn protobuf:compile-custom

mvn spring-boot:run

```

## build
```
projects/build.sh
```

## run
```sh
docker-compose -f projects/docker/docker-compose.yml up
```

### k8s
```sh
helm upgrade --install ticket-app k8s/helm/apps/ticket-app/  --set global.resourceRequest.enabled=false --set global.docker.repoHost=ethansmart-docker.pkg.coding.net/istioalltime
```
