FROM alpine:3.10
MAINTAINER lico603

# Install base packages
RUN apk update && apk add curl bash tree tzdata \
    && cp -r -f /usr/share/zoneinfo/Hongkong /etc/localtime

ENV  GIN_MODE=release

COPY lico_alone /
COPY www/ /www

RUN chmod u+x /lico_alone

EXPOSE 80

CMD ["/lico_alone"]
