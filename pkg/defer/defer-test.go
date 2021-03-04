package main

import (
	"fmt"
)

func deferCall() {
	fmt.Println("begin")
	defer func() { fmt.Println("AAA") }()
	defer func() { fmt.Println("BBB") }()
	defer func() { fmt.Println("CCC") }()
	fmt.Println("end")
}

func main() {
	deferCall()
}
