FROM golang:alpine AS budiler

ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

WORKDIR /MeiChat

COPY bin/connect_service_linux ./connect/connect_service_linux

VOLUME /MeiChat/conf

EXPOSE 8080

ENTRYPOINT ["/MeiChat/connect/connect_service_linux"]
