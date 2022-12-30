FROM golang:1.19 as builder

WORKDIR /go/workspace

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOPROXY="https://goproxy.cn,direct"

COPY go.mod go.sum ./

COPY . .

RUN go build test-server .

FROM alpine:3.9

ARG APP_PATH=/go/workspace
WORKDIR ${APP_PATH}

COPY --from=builder ${APP_PATH}/ ./

EXPOSE 8080

CMD ./test-server