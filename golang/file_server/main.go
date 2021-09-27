package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/prometheus/common/log"
)

func main() {
	dir := flag.String("dir", "/tmp", "默认路径")

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(*dir))
	mux.Handle("/", fs)
	fmt.Println("文件服务启动 文件路径为 ", dir)
	log.Fatalln(http.ListenAndServe(":10000", mux))
}

//func main() {
//	http.Handle("/", http.FileServer(http.Dir("./")))
//	http.ListenAndServe(":8080", nil)
//}
