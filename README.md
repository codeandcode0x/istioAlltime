# Service Mesh Alltime Project

## grpc
```sh
protoc -I ${protoDir}/ ${protoDir}/*proto --go_out=plugins=grpc:${outDir}
```
example:

Golang

```sh
protoc -I rpc/grpc/protos/movie/ rpc/grpc/protos/movie/*proto --go_out=plugins=grpc:rpc/grpc/protos/movie
```

Java

```sh
protoc --plugin=protoc-gen-grpc-java --grpc-java_out="$OUTPUT_FILE" --proto_path="$DIR_OF_PROTO_FILE" "$PROTO_FILE"

mvn protobuf:compile-custom

mvn spring-boot:run

```


# build
```sh
_deploy/build.sh
```

# run
```sh
docker-compose -f projects/docker/docker-compose.yml up
```

# build init data job
```sh
docker build -t roandocker/initdata-job:latest -f projects/docker/initdata-job/Dockerfile  projects/docker/initdata-job
```