FROM golang:1.15.5-alpine3.12 as goBuilder
ADD ./ /myapp
WORKDIR /myapp
RUN export CGO_ENABLED=0 && export GOPROXY=https://goproxy.io,direct && cd ./server && go build -v -o myapp

FROM alpine:3.12
MAINTAINER lico
RUN apk update && apk add curl bash tree tzdata \
    && cp -r -f /usr/share/zoneinfo/Hongkong /etc/localtime
COPY --from=goBuilder /myapp/server/myapp /
RUN chmod u+x /myapp
EXPOSE 80
CMD ["/myapp"]
