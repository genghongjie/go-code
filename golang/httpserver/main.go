package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var port *string

func main() {
	_ = os.Setenv("VERSION", "v1.00")
	port = flag.String("port", "80", "http server port")
	apiRun()
}

func apiRun() {
	httpServer := http.Server{
		Addr:    ":" + *port,
		Handler: initHandle(),
	}
	//启动服务
	go func() {
		log.Println("Server start, port is ", *port)
		log.Fatalln(httpServer.ListenAndServe())
	}()
	//优雅退出
	exitServer(httpServer)
}

//优雅退出
func exitServer(s http.Server) {
	// grace shutdown
	quit := make(chan os.Signal, 1)
	// receive system signal
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // block
	// service will be shut down in 5 seconds, wait for the request to be processed
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown server failed: %s", err)
	}
	log.Println("server shutdown successfully")
}

func initHandle() *http.ServeMux {
	//多路复用
	mux := http.NewServeMux()
	//健康检查
	mux.HandleFunc("/", myHeader)
	mux.HandleFunc("/health", myHandler)
	//性能分析模块
	performanceProfiling(mux)
	return mux
}

func myHeader(w http.ResponseWriter, r *http.Request) {
	for key := range r.Header {
		w.Header().Set(key, r.Header.Get(key))
	}
	w.Header().Set("version", os.Getenv("VERSION"))
	w.WriteHeader(http.StatusOK)
	//打印访问信息
	log.Printf("remote addr %s,  ip %s, http code %d, methpd  %s", r.RemoteAddr, r.Host, http.StatusOK, r.Method)

}

func myHandler(w http.ResponseWriter, r *http.Request) {
	//环境变量中查询版本
	if version, ok := os.LookupEnv("VERSION"); ok {
		w.Header().Set("version", version)
	}
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("It is ok!"))
	if err != nil {
		log.Println(err.Error())
	}
	//打印访问信息
	log.Printf("remote addr %s,  ip %s, http code %d, methpd  %s", r.RemoteAddr, r.Host, http.StatusOK, r.Method)
}

//性能分析
func performanceProfiling(mux *http.ServeMux) {
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
}
