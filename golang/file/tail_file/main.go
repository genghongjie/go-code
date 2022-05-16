package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func main() {

	var fileName string
	flag.StringVar(&fileName, "f", "", "采集数据文件路径")
	flag.Parse()
	if fileName == "" {
		panic("f 参数不能为空，需要输入监听的文件路径名")
	}
	t, err := tail.TailFile(fileName, tail.Config{Follow: true, Logger: tail.DiscardingLogger})

	if err != nil {
		panic(err)
	}
	for line := range t.Lines {
		fmt.Println(fmt.Sprintf("Time: %s Line: %s", time.Now().Format("2006-01-02 15:04:05"), line.Text))
	}
}
