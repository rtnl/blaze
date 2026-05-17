FROM alpine:latest

WORKDIR /opt/rtnl/blaze

RUN apk add gcompat

ADD ./build/container /opt/rtnl/blaze/

ENTRYPOINT [ "./api" ]
