FROM alpine

MAINTAINER tamura2004@gmail.com

RUN apk update \
	&& apk add ruby \
	&& apk add curl \
	&& rm -rf /var/cache/apk/*

WORKDIR /home/project/ruby
COPY server.rb .
COPY client.rb .
COPY badclient/badclient /usr/local/bin/
COPY badserver/badserver /usr/local/bin/

CMD ["badserver"]