package main

import (
	"log"
	"net/http"
)

func main() {

	myHandler := func(w http.ResponseWriter, r *http.Request) {
		//读取header信息
		token := r.Header.Get("toke")

		//response中写入header
		w.Header().Set("token", token)
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("It is ok!"))
		if err != nil {
			log.Fatalln(err.Error())
		}
		//打印访问信息
		log.Printf("%s  %d  %s", r.RemoteAddr, http.StatusOK, r.Method)
	}

	//健康检查
	http.HandleFunc("/healthz", myHandler)

	//启动服务
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
