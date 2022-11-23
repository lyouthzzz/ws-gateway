FROM golang:1.16-alpine3.12 AS build

ENV GOPROXY=https://goproxy.cn,direct
ENV GOPRIVATE=*.weimob.com
ENV GO111MODULE=auto
ENV CGO_ENABLED=0
ENV GOOS=linux

ARG APP_NAME
ARG LDFLAGS="-s -w -linkmode external -extldflags \"-static\""

WORKDIR /go/src/ws-gateway

RUN sed -i 's!http://dl-cdn.alpinelinux.org/!https://mirrors.tuna.tsinghua.edu.cn/!g' /etc/apk/repositories
RUN apk update && apk upgrade
RUN apk add --no-cache git openssh gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -installsuffix=cgo \
    -tags="jsoniter netgo" \
    -ldflags="$LDFLAGS" \
    -o=${APP_NAME} \
    /go/src/ws-gateway/app/${APP_NAME}/cmd/${APP_NAME}

FROM alpine:3.12

ARG APP_NAME

ENV APP_NAME ${APP_NAME}
ENV TZ=Asia/Shanghai LANG=C.UTF-8 GOPATH=/go

RUN sed -i 's!http://dl-cdn.alpinelinux.org/!https://mirrors.tuna.tsinghua.edu.cn/!g' /etc/apk/repositories

RUN apk --update add ca-certificates tzdata && \
    cp /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone

WORKDIR /app

COPY --from=build /go/src/ws-gateway/${APP_NAME} ${APP_NAME}

CMD /app/${APP_NAME}
