package main

import "fmt"

func main() {
	type People struct {
		Name string
		Age  int
	}
	hank := People{Name: "Hank", Age: 20}
	fmt.Printf("%v \n", hank)
	fmt.Printf("%+v \n", hank)
	fmt.Printf("%#v \n", hank)
	fmt.Printf("%T \n", hank.Name)

}
