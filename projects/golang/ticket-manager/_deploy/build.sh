#!/bin/bash

set -eo pipefail

appVersion=$1

if [ ! $appVersion ]; then
  appVersion="latest" 
fi

docker build -t roandocker/ticket-manager:$appVersion -f _deploy/Dockerfile  --no-cache .

sleep 10 && docker system prune -f

docker push roandocker/ticket-manager:$appVersion

export TICKET_VERSION=$appVersion