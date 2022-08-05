package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "1|2|3||4"

	s := strings.Split(str, "|")

	fmt.Println(len(s))
	fmt.Println(s[3])
	fmt.Println(len(s))
}
