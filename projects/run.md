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