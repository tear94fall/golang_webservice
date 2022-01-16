#!/bin/bash

docker stop go-web
docker rm go-web
docker rmi golang-webservice

#./build.sh
#./run.sh
