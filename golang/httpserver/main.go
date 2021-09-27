package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	myHandler := func(w http.ResponseWriter, r *http.Request) {
		//读取header信息
		token := r.Header.Get("toke")

		//环境变量中查询版本
		if version, ok := os.LookupEnv("V"); ok {
			w.Header().Set("version", version)
		}

		//response中写入header
		w.Header().Set("token", token)
		w.Header().Set("remote_addr", r.RemoteAddr)
		w.Header().Set("method", r.Method)
		w.Header().Set("host", r.Host)
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("It is ok!"))
		if err != nil {
			log.Fatalln(err.Error())
		}

		//打印访问信息
		log.Printf("%s  %s  %d  %s", r.RemoteAddr, r.Host, http.StatusOK, r.Method)
	}

	//健康检查
	http.HandleFunc("/healthz", myHandler)

	_ = os.Setenv("V", "v1.00")
	//启动服务
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
