package main

import "fmt"

func fuck(c chan int) {
	data := <-c
	fmt.Println(data)
}

func main() {
	c := make(chan int)
	go fuck(c)
	c <- 12
}
