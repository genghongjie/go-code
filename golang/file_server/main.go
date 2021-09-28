package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"runtime/pprof"
)

func main() {
	dir := flag.String("dir", "./", "默认路径")
	cpuProfile := flag.String("cpu_profile", "", "cpu 分析写入到指定文件，可通过 go tool prop /tmp/cpufile 进行分析")

	flag.Parse()

	//cpu性能分析模块
	performanceProfiling(cpuProfile)

	fs := http.FileServer(http.Dir(*dir))

	mux := http.NewServeMux()
	mux.Handle("/", fs)

	log.Println("文件服务启动 文件路径为 ", *dir)
	log.Fatalln(http.ListenAndServe(":10000", mux))
}

//性能分析开关
func performanceProfiling(cpuProfile *string) {
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal("文件", *cpuProfile, "创建失败,", err.Error())
		}
		_ = pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
}
