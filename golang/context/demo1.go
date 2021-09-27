package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		fmt.Println("worker")
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("worker kill")
			break
		default:
		}

	}
}

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	defer cancel()
	go worker(ctx)

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

	fmt.Println("worker over")
}
