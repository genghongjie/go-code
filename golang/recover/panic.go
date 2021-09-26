package main

import "fmt"

func main() {
	fmt.Println("this one")
	testPanic()
	fmt.Println("this two")
}

func testPanic() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	panic("now is panic")
}
