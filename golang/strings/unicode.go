package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {

	str := "Go爱好者"
	fmt.Printf("The string: %q\n", str)
	fmt.Printf("  => runes(char): %q\n", []rune(str))
	fmt.Printf("  => runes(hex): %x\n", []rune(str))
	fmt.Printf("  => bytes(hex): [% x]\n", []byte(str))
	fmt.Println([]rune(str))

	s1 := strconv.QuoteToASCII("苏苏")
	fmt.Println(s1)
	fmt.Println(strconv.Unquote(s1))

	s := "Hello 我的世界 go！"
	for _, r := range s {
		// 判断字符是否为汉字
		if unicode.Is(unicode.Scripts["Han"], r) {
			fmt.Printf("%c\n", r)
		}
	}

	for _, r1 := range s {
		if !unicode.IsControl(r1) {
			fmt.Println(r1)
		}
	}
}
