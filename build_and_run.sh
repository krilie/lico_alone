#!/bin/bash -x

export PATH=$PATH:/opt/go/bin/

go build .

docker build -t lico603/lico_user:$BUILD_NUMBER .

docker stop lico_user
docker rm lico_user
docker run --name lico_user -p 1000:8080 lico603/lico_user:$BUILD_NUMBER

echo "end"

