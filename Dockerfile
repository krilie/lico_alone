FROM golang:1.14.4-alpine3.12 as goBuilder
ADD ./ /myapp
WORKDIR /myapp
RUN export CGO_ENABLED=0 && export GOPROXY=https://goproxy.io,direct && cd ./server && go build -v -o myapp

FROM node:lts-alpine3.10 as webBuilder
ADD ./ /myapp
WORKDIR /myapp
RUN cd ./web && npm build

FROM alpine:3.11
MAINTAINER livo
RUN apk update && apk add curl bash tree tzdata \
    && cp -r -f /usr/share/zoneinfo/Hongkong /etc/localtime
COPY --from=goBuilder /myapp/server/elflog /
COPY --from=webBuilder /myapp/web/dist /www
RUN chmod u+x /elflog
EXPOSE 80
CMD ["/elflog"]
