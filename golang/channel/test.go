package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	//生产
	go Prod(ch)
	//消费
	go Consumer(ch)
	go Consumer(ch)
	//等待5分钟
	done := time.After(5 * time.Minute)
	<-done
}
func Prod(in chan<- int) {
	sum := 0
	done := time.Tick(time.Second)
	for {
		sum++
		select {
		case <-done:
			fmt.Println("prod: ", sum)
			in <- sum
		}
	}
}
func Consumer(out <-chan int) {
	for outItem := range out {
		time.Sleep(3 * time.Second)
		fmt.Println("consumer: ", outItem)
	}
}
