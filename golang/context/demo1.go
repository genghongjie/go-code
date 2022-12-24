package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		fmt.Println("worker ing")
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("worker kill")
			break
		default:
			fmt.Println("working working working ")
		}

	}
}

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(4*time.Second))
	defer cancel()
	go worker(ctx)

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("ctx.Err()")
		fmt.Println(ctx.Err())
	}

	fmt.Println("worker over")
}
