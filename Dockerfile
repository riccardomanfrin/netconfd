# syntax=docker/dockerfile:experimental
FROM golang:1.12.9-alpine3.10 as builder
ARG CI_PROJECT_PATH
RUN mkdir -p -m 0600 ~/.ssh
RUN echo "StrictHostKeyChecking no" >> ~/.ssh/config
RUN apk update && \
    apk add --no-cache git make openssh pkgconfig gcc
COPY ./Athonet_Services_CA.crt /usr/local/share/ca-certificates/Athonet_Services_CA.crt
RUN update-ca-certificates
ENV GOPATH /go/
ENV SRCPATH $GOPATH/src/gitlab.lan.athonet.com/$CI_PROJECT_PATH
ENV STATIKPATH $SRCPATH/vendor/github.com/rakyll/statik
ADD src.tar.gz $SRCPATH
RUN --mount=type=ssh \
	cd $SRCPATH && \
    make

FROM alpine:3.7
ARG CI_PROJECT_PATH
ARG CI_PROJECT_NAME
ENV GOPATH /go/
ENV SRCPATH $GOPATH/src/gitlab.lan.athonet.com/$CI_PROJECT_PATH
RUN mkdir /app/
COPY --from=builder $SRCPATH/$CI_PROJECT_NAME /app/$CI_PROJECT_NAME
RUN mkdir /conf/
COPY --from=builder $SRCPATH/$CI_PROJECT_NAME.json /conf/
EXPOSE 10000/tcp
EXPOSE 10002/tcp
WORKDIR /app/
ENTRYPOINT ["./$CI_PROJECT_NAME", "--config", "/conf/$CI_PROJECT_NAME.json"]