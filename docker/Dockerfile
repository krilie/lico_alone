FROM alpine:3.11
MAINTAINER livo

# Install base packages
RUN apk update && apk add curl bash tree tzdata \
    && cp -r -f /usr/share/zoneinfo/Hongkong /etc/localtime

COPY lico_alone /
COPY www/ /www

RUN chmod u+x /lico_alone

EXPOSE 80
CMD ["/lico_alone"]
