FROM golang:1.17.2-alpine AS build
WORKDIR /go/src/genghongjie/httpserver
COPY . /go/src/genghongjie/httpserver/
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app-lib main.go

FROM alpine
WORKDIR /app
COPY --from=build /go/src/genghongjie/httpserver/app-lib  /app/app-lib

ENV TZ=Asia/Shanghai
RUN  export LANG=zh_CN.UTF-8

CMD [ "./app-lib"]
EXPOSE 8080
