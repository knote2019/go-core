package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {

	var dataChannel = make(chan bool, 20000)
	var beginChannel = make(chan bool)

	go func() {
		runtime.LockOSThread()
		<-beginChannel
		fmt.Println("beginChannel")
		tm := time.Now()
		for i := 0; i < 10000000; i++ {
			<-dataChannel
		}
		fmt.Println(time.Now().Sub(tm))
		os.Exit(0)
	}()

	for i := 0; i < 50000; i++ {
		// 负载
		go func() {
			var count int
			load := 100000
			for {
				count++
				if count >= load {
					count = 0
					runtime.Gosched()
				}
			}
		}()
	}

	for i := 0; i < 20; i++ {
		go func() {
			for {
				dataChannel<- true
			}
		}()
	}

	fmt.Println("all start")
	beginChannel<- true

}
