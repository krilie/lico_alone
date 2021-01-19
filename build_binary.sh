#!/bin/bash

docker run -it -v "$PWD":/app -v "$PWD"/.cache/gopath:/root/go -v "$PWD"/.cache/gocache:/root/.cache/go-build golang:1.15.5-alpine3.12 \
           /bin/sh -c "cd /app && export CGO_ENABLED=0 && export GOPROXY=https://goproxy.io,direct && cd ./server && go build -v -o myapp"
if [ "$?" -eq 0 ];then
  echo "build ok"
else
  echo "no"
  exit 1
fi
