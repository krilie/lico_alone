FROM alpine:3.9
MAINTAINER lico603

COPY ./lico-my-site-user /opt/lico_user
COPY ./enterpoint.sh /opt/enterpoint.sh
RUN chmod u+x /opt/enterpoint.sh \
    &&chmod u+x /opt/lico_user

EXPOSE 8080

ENTRYPOINT ["/opt/enterpoint.sh"]
