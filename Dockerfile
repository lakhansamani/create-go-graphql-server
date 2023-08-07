FROM alpine:latest

ARG TARGETPLATFORM
RUN apk add -u ca-certificates
RUN echo ${TARGETPLATFORM}
ADD ./bin/${TARGETPLATFORM}/app /app/
RUN mkdir -p /etc/app

WORKDIR /app/
ENTRYPOINT [ "/app/app" ]