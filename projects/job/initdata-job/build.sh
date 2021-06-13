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
	repoUser=$2
	docker build -t roandocker/initdata-job:$version -f projects/docker/initdata-job/Dockerfile --no-cache  projects/docker/initdata-job
}

# prune image
function prune() {
	sleep 10 && docker system prune -f
}

# push image
function push() {
	version=$1
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