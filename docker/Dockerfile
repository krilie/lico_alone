FROM alpine:3.12
MAINTAINER lico
# Install base packages
RUN apk update && apk add curl bash tree tzdata \
    && cp -r -f /usr/share/zoneinfo/Hongkong /etc/localtime
COPY lico_alone /
COPY migrations /migrations
RUN chmod u+x /lico_alone
EXPOSE 80
CMD ["/lico_alone"]
