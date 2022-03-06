#!/bin/bash

# stop container
docker stop go-web

# remove container
docker rm go-web

# remove container images
docker rmi golang-webservice

# build container
docker build -t golang-webservice .

# run container
docker run --name go-web -d -it -p 8080:8080 -v `pwd`/tmp:/app/attaches golang-webservice