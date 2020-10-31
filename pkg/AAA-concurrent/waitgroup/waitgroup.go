package main

import "fmt"

func main() {
	done := make(chan int, 10) // 带 10 个缓存

	// 开N个后台打印线程
	for i := 0; i < cap(done); i++ {
		go func(){
			fmt.Println("hello kenny")
			done <- 1
		}()
	}

	// 等待N个后台线程完成
	for i := 0; i < cap(done); i++ {
		<-done
	}
}

// wg.Add(1)用于增加等待事件的个数
// wg.Done()表示完成一个事件
