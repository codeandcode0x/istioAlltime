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
	docker build -t $repoUser/ticket-manager:$version -f projects/golang/ticket-manager/_deploy/Dockerfile  --no-cache .
	docker build -t $repoUser/initdata-job:$version -f docker/initdata-job/Dockerfile  --no-cache .
}

# prune image
function prune() {
	sleep 10 && docker system prune -f
}

# push image
function push() {
	version=$1
	docker push $repoUser/ticket-manager:$version
	docker push $repoUser/initdata-job:$version
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