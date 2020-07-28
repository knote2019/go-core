package main

import (
	"fmt"
)

func defer_call() {
	defer func() { fmt.Println("AAA") }()
	defer func() { fmt.Println("BBB") }()
	defer func() { fmt.Println("CCC") }()
	panic("kenny exception")
	fmt.Println("test")
}

func main() {
	defer_call()
}
