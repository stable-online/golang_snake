#!/bin/bash

docker ps -a | grep snake | awk '{print $1}' | xargs docker stop | xargs docker rm
docker images | grep snake | awk '{print $3}' | xargs docker rmi

docker build -t snake:v1 .