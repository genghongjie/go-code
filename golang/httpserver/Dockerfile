FROM golang:1.17.0-alpine AS build
WORKDIR /go/src/genghongjie/httpserver
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app-lib main.go

FROM alpine:3.13.6
WORKDIR /app
ENV TZ=Asia/Shanghai
RUN  export LANG=zh_CN.UTF-8

COPY --from=build /go/src/genghongjie/httpserver/app-lib  /app/app-lib
CMD [ "./app-lib"]
EXPOSE 80