package main

import (
	"fmt"
	"time"
)

func fuck(dataChannel chan int) {
	data := <-dataChannel
	fmt.Println(data)
}

func main() {
	dataChannel := make(chan int)
	go fuck(dataChannel)
	dataChannel <- 12
	time.Sleep(time.Second)
}
