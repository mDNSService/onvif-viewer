FROM alpine
LABEL name=onvif-viewer
LABEL url=https://github.com/OpenIoTHub/mDNSService
RUN apk add --no-cache bash

WORKDIR /app
COPY onvif-viewer /app/
ENV TZ=Asia/Shanghai
#mdns端口
EXPOSE 5353/udp
EXPOSE 34324
ENTRYPOINT ["/app/onvif-viewer"]
CMD ["-c", "/root/config.yaml"]