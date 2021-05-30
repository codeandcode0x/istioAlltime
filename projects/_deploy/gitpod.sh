#!/bin/bash

set -e

sudo apt-get install docker-compose -y

docker-compose -f docker-compose.yml -p ticket  up
