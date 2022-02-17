package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"runtime/pprof"
)

func main() {
	dir := flag.String("dir", "./", "默认路径")
	port := flag.String("port", "10000", "启动端口")
	cpuProfile := flag.String("cpu_profile", "", "cpu 分析写入到指定文件，可通过 go tool prop /tmp/cpufile 进行分析")

	flag.Parse()

	//cpu性能分析模块
	performanceProfiling(cpuProfile)

	fs := http.FileServer(http.Dir(*dir))

	mux := http.NewServeMux()
	mux.Handle("/", fs)

	log.Println("文件服务启动 文件路径为 ", *dir)
	log.Println("访问路径微 http://127.0.0.1:" + *port)

	ip, err := externalIP()
	if err != nil {
		fmt.Println(err)
	}
	log.Println("访问路径微 http://" + ip.String() + ":" + *port)
	log.Fatalln(http.ListenAndServe(":"+*port, mux))
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

func externalIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network?")
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}
