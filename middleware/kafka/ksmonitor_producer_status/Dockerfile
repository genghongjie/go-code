FROM golang:1.18 AS build
WORKDIR /go/src/github.com/genghongjie/go-code/middleware/kafka/ksmonitor_producer_status
COPY . .
ENV GOPROXY="https://proxy.golang.com.cn,direct"
RUN go mod tidy
RUN go build -o app-lib
RUN chmod u+x app-lib
#
FROM 192.168.102.57:5000/library/alpine-glibc:latest
WORKDIR /app
RUN  export LANG=zh_CN.UTF-8

COPY --from=build  /go/src/github.com/genghongjie/go-code/middleware/kafka/ksmonitor_producer_status/app-lib  /app/app-lib


RUN chmod u+x app-lib
CMD [ "/app/app-lib"]