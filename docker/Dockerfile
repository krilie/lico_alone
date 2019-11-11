FROM alpine:3.10
MAINTAINER lico603

# Install base packages
RUN apk update && apk add curl bash tree tzdata \
    && cp -r -f /usr/share/zoneinfo/Hongkong /etc/localtime

ENV  GIN_MODE=release

COPY lico_alone enterpoint.sh /opt/

RUN chmod u+x /opt/enterpoint.sh \
    && chmod u+x /opt/lico_alone

EXPOSE 80

ENTRYPOINT ["/opt/enterpoint.sh"]
