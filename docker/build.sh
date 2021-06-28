#!/bin/bash

set -eo pipefail

appVersion=$1
repoUser=$2
repoHost=$3

if [ ! $appVersion ]; then
  appVersion="latest" 
fi

if [ ! $repoUser ]; then
  repoUser="roandocker" 
fi

if [ ! $repoHost ]; then
  repoHost="docker.io" 
fi

# build image
function build() {
	version=$1
	repoUser=$2
	repoHost=$3
	docker build -t $repoHost/$repoUser/ticket-manager:$version -f docker/projects/golang/ticket-manager/Dockerfile  --no-cache ./
	docker build -t $repoHost/$repoUser/ticket-frontend:$version -f docker/projects/java/ticket-frontend/Dockerfile  --no-cache ./
	docker build -t $repoHost/$repoUser/initdata-job:$version -f docker/projects/job/initdata-job/Dockerfile --no-cache  projects/job/initdata-job
	docker build -t $repoHost/$repoUser/k8s-wait-for:$version -f docker/tools/k8s-wait-for/Dockerfile --no-cache  docker/tools/k8s-wait-for
}

# prune image
function prune() {
	sleep 10 && docker system prune -f
}

# push image
function push() {
	version=$1
	repoUser=$2
	repoHost=$3
	docker push $repoHost/$repoUser/ticket-manager:$version
	docker push $repoHost/$repoUser/ticket-frontend:$version
	docker push $repoHost/$repoUser/initdata-job:$version
	docker push $repoHost/$repoUser/k8s-wait-for:$version
}



#main 
function main() {
	build $1 $2 $3
	prune
	push $1 $2 $3
}

# run main
main $appVersion $repoUser $repoHost

export TICKET_VERSION=$appVersion