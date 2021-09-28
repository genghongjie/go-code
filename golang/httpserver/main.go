package main

import (
	"log"
	"net/http"
	"net/http/pprof"
	"os"
)

func main() {

	//多路复用
	mux := http.NewServeMux()
	//健康检查
	mux.HandleFunc("/healthz", myHandler)
	//性能分析模块
	performanceProfiling(mux)

	_ = os.Setenv("V", "v1.00")
	//启动服务
	log.Fatalln(http.ListenAndServe(":8080", mux))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
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
		log.Println(err.Error())
	}
	//打印访问信息
	log.Printf("%s  %s  %d  %s", r.RemoteAddr, r.Host, http.StatusOK, r.Method)
}

func performanceProfiling(mux *http.ServeMux) {
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
}
