package main

import (
	"fmt"
	"time"
)

func c1Sender(ch chan string) {
	for true {
		ch <- "channel1"
		time.Sleep(time.Second)
	}
}

func c2Sender(ch chan string) {
	for true {
		ch <- "channel2"
		time.Sleep(time.Second)
	}
}

func c3Sender(ch chan string) {
	for true {
		ch <- "channel3"
		time.Sleep(time.Second)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go c1Sender(c1)
	go c2Sender(c2)
	go c3Sender(c3)

	for true {
		select {
		case v1 := <-c1:
			fmt.Println(v1)
		case v2 := <-c2:
			fmt.Println(v2)
		case v3 := <-c3:
			fmt.Println(v3)
		}
	}
}
