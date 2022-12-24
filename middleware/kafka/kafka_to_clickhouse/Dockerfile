FROM jit.dockerhub.com:5000/basic/showcase-admin:latest AS build
WORKDIR /go/src/github.com/genghongjie/go-code/middleware/kafka
COPY . .
RUN go build -o app-lib producer.go
RUN chmod u+x app-lib
#
FROM jit.dockerhub.com:5000/alpine-glibc:latest
WORKDIR /app
RUN  export LANG=zh_CN.UTF-8

COPY --from=build  /go/src/github.com/genghongjie/go-code/middleware/kafka/app-lib  /app/app-lib
COPY --from=build  /go/src/github.com/genghongjie/go-code/middleware/kafka/getting-started.properties  /app/getting-started.properties

RUN chmod u+x app-lib
CMD [ "/app/app-lib ","getting-started.properties"]