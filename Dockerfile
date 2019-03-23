FROM alpine:3.9
MAINTAINER lico603

COPY ./lico-my-site-user /opt/lico_user
COPY ./enterpoint.sh /opt/enterpoint.sh

EXPOSE 8080

ENTRYPOINT ["/opt/enterpoint.sh"]
