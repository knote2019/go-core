package main

import "fmt"

func main() {
	dataChannel := make(chan int, 10) // 带 10 个缓存

	// 开N个后台打印线程
	for i := 0; i < cap(dataChannel); i++ {
		go func(){
			fmt.Println("hello kenny")
			dataChannel <- 1
		}()
	}

	// 等待N个后台线程完成
	for i := 0; i < cap(dataChannel); i++ {
		<-dataChannel
	}
}

// wg.Add(1)用于增加等待事件的个数
// wg.dataChannel()表示完成一个事件
