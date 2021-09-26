package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	var age int
	flag.StringVar(&name, "name", "Hank", "用户名")
	flag.IntVar(&age, "age", 19, "年龄")
	flag.Parse()
	flag.Visit(getFlagValue)

}
func getFlagValue(f *flag.Flag) {
	fmt.Println(f.Name, ":", f.Value.String())
}
