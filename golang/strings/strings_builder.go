package main

import (
	"fmt"
	"strings"
)

func main() {

	//strings.Builder类型的值（以下简称Builder值）的优势有下面的三种：
	// 已存在的内容不可变，但可以拼接更多的内容；
	// 减少了内存分配和内容拷贝的次数；
	// 可将内容重置，可重用值。

	var build strings.Builder
	//手动扩容
	build.Grow(20)
	fmt.Printf(">> build 初始长度 %d \n", build.Len())
	fmt.Printf(">> build 初始容量  %d \n", build.Cap())

	fmt.Println("第一次追加内容：")
	build.WriteString("My name is hank")
	fmt.Println(build.String())
	fmt.Printf(">> build 内容长度 %d \n", build.Len())
	fmt.Printf(">> build 容量  %d \n", build.Cap())

	//自动扩容
	fmt.Println("第二次追加内容：")
	build.WriteString(" and my age is 20")
	fmt.Println(build.String())
	fmt.Printf(">> build 内容长度 %d \n", build.Len())
	fmt.Printf(">> build 容量  %d \n", build.Cap())

	name := "我的英文名字是 Hank.geng"
	//判断名字中是否包含字符串 Hank
	isExist := strings.Contains(name, "Hank")
	fmt.Printf("字符串中是否包含Hank : %t", isExist)

}
