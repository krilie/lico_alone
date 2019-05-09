FROM alpine:3.9
MAINTAINER lico603

ENV  GIN_MODE=release
COPY ./lico_alone /opt/app
COPY enterpoint.sh /opt/enterpoint.sh

RUN chmod u+x /opt/enterpoint.sh \
    && chmod u+x /opt/app

EXPOSE 443

ENTRYPOINT ["/opt/enterpoint.sh"]
