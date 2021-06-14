#!/bin/bash

set -eo pipefail

appVersion=$1

if [ ! $appVersion ]; then
  appVersion="latest" 
fi

repoUser="roandocker"

# build image
function build() {
	version=$1
	docker build -t $repoUser/ticket-manager:$version -f docker/projects/golang/ticket-manager/Dockerfile  --no-cache projects/golang/ticket-manager/
	docker build -t $repoUser/ticket-frontend:$version -f docker/projects/java/ticket-frontend/Dockerfile  --no-cache projects/java/ticket-frontend/
	docker build -t $repoUser/initdata-job:$version -f docker/projects/job/initdata-job/Dockerfile --no-cache  projects/job/initdata-job
	docker build -t $repoUser/k8s-wait-for:$version -f docker/tools/k8s-wait-for/Dockerfile --no-cache  docker/tools/k8s-wait-for
}

# prune image
function prune() {
	sleep 10 && docker system prune -f
}

# push image
function push() {
	version=$1
	docker push $repoUser/ticket-manager:$version
	docker push $repoUser/ticket-frontend:$version
	docker push $repoUser/initdata-job:$version
	docker push $repoUser/k8s-wait-for:$version
}



#main 
function main() {
	build $1 $2
	prune
	push $1 $2
}

# run main
main $appVersion $repoUser

export TICKET_VERSION=$appVersion