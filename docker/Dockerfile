FROM alpine:3.9
MAINTAINER lico603

ENV  GIN_MODE=release

COPY lico_alone /opt/lico_alone
COPY enterpoint.sh /opt/enterpoint.sh

RUN chmod u+x /opt/enterpoint.sh \
    && chmod u+x /opt/lico_alone

EXPOSE 80

ENTRYPOINT ["/opt/enterpoint.sh"]
