FROM alpine:3.7

LABEL maintainer="FlowQ proxy Docker Maintainers <steven@flowq.io>"


COPY ./proxy /proxy
EXPOSE 8899

CMD ["/proxy","-h", "0.0.0.0"]

