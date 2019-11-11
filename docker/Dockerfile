FROM alpine:3.10
MAINTAINER lico603

ENV  GIN_MODE=release

COPY lico_alone enterpoint.sh /opt/

RUN chmod u+x /opt/enterpoint.sh \
    && chmod u+x /opt/lico_alone \
    && cd opt && ls -alh

EXPOSE 80

ENTRYPOINT ["/bin/sh"]
