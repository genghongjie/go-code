package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	dir := flag.String("dir", "./", "默认路径")

	flag.Parse()
	fs := http.FileServer(http.Dir(*dir))

	mux := http.NewServeMux()
	mux.Handle("/", fs)

	log.Println("文件服务启动 文件路径为 ", *dir)
	log.Fatalln(http.ListenAndServe(":10000", mux))
}
