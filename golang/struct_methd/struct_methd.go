package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Persion struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	p := Persion{"hank", 1000}

	testType := reflect.TypeOf(p)
	f := testType.Field(0)
	fmt.Printf("%v\n", f)
}
