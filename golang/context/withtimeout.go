package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	done := make(chan int)
	go func(ctx context.Context) {

		for i := 0; i < 2; i++ {
			fmt.Println("111")
			time.Sleep(1 * time.Second)
		}
		done <- 1
	}(ctx)

	select {
	case <-done:
		fmt.Println("done")
		return
	//case <-time.After(1 * time.Second):
	//	fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}
