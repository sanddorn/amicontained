FROM golang:alpine as builder
MAINTAINER Nils Bokermann <nils.bokermann@bermuda.de>

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

RUN	apk add --no-cache \
	bash \
	ca-certificates

COPY . /go/src/github.com/sanddorn/checkcapabilities

RUN set -x \
	&& apk add --no-cache --virtual .build-deps \
		git \
		gcc \
		libc-dev \
		libgcc \
		make \
	&& cd /go/src/github.com/sanddorn/checkcapabilities \
	&& make static \
	&& mv checkcapabilities /usr/bin/checkcapabilities \
	&& apk del .build-deps \
	&& rm -rf /go \
	&& echo "Build complete."

FROM alpine:latest

COPY --from=builder /usr/bin/checkcapabilities /usr/bin/checkcapabilities
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs

ENTRYPOINT [ "checkcapabilities" ]
CMD [ "--help" ]
