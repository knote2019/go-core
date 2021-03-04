package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex

	mu.Lock()
	go func() {
		fmt.Println("hello")
		mu.Unlock()
	}()

	mu.Lock() // 会阻塞在此

	time.Sleep(time.Second)
	fmt.Println("end")
}
