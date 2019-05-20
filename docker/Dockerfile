FROM alpine:3.9
MAINTAINER lico603

ENV  GIN_MODE=release \
  LICO_SERVICE_PORT=443 \
  LICO_DB_IP=local_mysql \
  LICO_DB_PORT=3306 \
  LICO_DB_USER=root \
  LICO_DB_PASSWORD=123456#b*a \
  LICO_SSL_PUBLIC_KEY=/cert/full_chain.pem \
  LICO_SSL_PRIVATE_KEY=/cert/private.key \
  LICO_OSS_KEY=23123 \
  LICO_OSS_SECRET=412rwer2342S \
  LICO_OSS_ENDPOINT=https://local_minio:9000 \
  LICO_OSS_BUCKET=image

COPY certs/ /cert/
COPY lico_alone enterpoint.sh /opt/

RUN chmod u+x /opt/enterpoint.sh \
    && chmod u+x /opt/lico_alone

EXPOSE 443

ENTRYPOINT ["/opt/enterpoint.sh"]
