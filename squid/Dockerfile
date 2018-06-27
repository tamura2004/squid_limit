FROM alpine

MAINTAINER tamura2004@gmail.com

RUN apk update \
	&& apk add squid \
	&& rm -rf /var/cache/apk/*

COPY start-squid.sh /usr/local/bin/

CMD ["/usr/local/bin/start-squid.sh"]