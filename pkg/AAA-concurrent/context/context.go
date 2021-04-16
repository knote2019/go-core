package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, number int) error {
	defer wg.Done()
	fmt.Println(number)
	for {
		select {
		default:
			fmt.Println("hello")
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go worker(ctx, &wg, i)
	}

	time.Sleep(time.Second)
	cancel()

	wg.Wait()
}
