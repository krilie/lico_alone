FROM alpine:3.9
MAINTAINER lico603

ENV  GIN_MODE=release
ENV  LICO_SERVICE_PORT=80
ENV  LICO_DB_IP=ligo.ml
ENV  LICO_DB_PORT=3306
ENV  LICO_DB_USER=root
ENV  LICO_DB_PASSWORD=123456#b*a
ENV  LICO_SSL_PUBLIC_KEY=/cert/full_chain.pem
ENV  LICO_SSL_PRIVATE_KEY=/cert/private.key

COPY certs/ /cert/
COPY lico_alone /opt/lico_alone
COPY enterpoint.sh /opt/enterpoint.sh

RUN chmod u+x /opt/enterpoint.sh \
    && chmod u+x /opt/lico_alone

EXPOSE 80

ENTRYPOINT ["/opt/enterpoint.sh"]
