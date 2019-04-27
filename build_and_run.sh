#!/bin/bash -x

export PATH=$PATH:/opt/go/bin/
export CGO_ENABLED=0

go build .

mv ./lico_alone ./docker/
cd ./docker

#lico603为dockerhub帐号
docker build -t lico603/lico_user:$BUILD_NUMBER .

docker stop lico_alone
docker rm lico_alone
docker run --name lico_alone -d -p 443:443 lico603/lico_user:$BUILD_NUMBER

echo "end"

