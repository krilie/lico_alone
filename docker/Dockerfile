FROM alpine:3.9
MAINTAINER lico603

ENV  GIN_MODE=release \
  LICO_SERVICE_PORT=80 \
  LICO_DB_IP=ligo.ml \
  LICO_DB_PORT=3306 \
  LICO_DB_USER=root \
  LICO_DB_PASSWORD=123456#b*a \
  LICO_SSL_PUBLIC_KEY=/opt/cert/full_chain.pem \
  LICO_SSL_PRIVATE_KEY=/opt/cert/private.key

COPY certs/ lico_alone enterpoint.sh /opt/

RUN chmod u+x /opt/enterpoint.sh \
    && chmod u+x /opt/lico_alone

EXPOSE 80

ENTRYPOINT ["/opt/enterpoint.sh"]
