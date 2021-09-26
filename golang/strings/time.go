package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format("2006 01 02 15:04:05"))

	var s = 1000
	sTime := time.Duration(s) * time.Second

	fmt.Println(sTime.String())

	//c := make(chan int)
	//go func() {
	//	time.Sleep(11 * time.Second)
	//	c <- 1000
	//}()
	//select {
	//case m := <-c:
	//	fmt.Println("接收值为 ", m)
	//case <-time.After(10 * time.Second):
	//	fmt.Println("10 秒内没有接收到消息，timed out")
	//}
	//
	//time.Sleep(100 * time.Microsecond)

	//tick
	cc := time.Tick(5 * time.Second)
	index := 0
	for now := range cc {
		index++
		fmt.Println(index)
		time.Sleep(7 * time.Second)
		fmt.Printf("%v\n", now.Second())
	}

	for {
		select {
		case now := <-cc:
			fmt.Printf("%v %s\n", now, "1")
		}
	}
}
