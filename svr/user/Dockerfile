FROM golang:alpine AS budiler

ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

WORKDIR /MeiChat

COPY bin/user_service_linux ./user/user_service_linux

VOLUME /MeiChat/conf

EXPOSE 9000

ENTRYPOINT ["/MeiChat/user/user_service_linux"]

