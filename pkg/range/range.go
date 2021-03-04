package main

import "fmt"

func main() {
	a := "kenny"
	for i, v := range a {
		fmt.Println(i, v)
	}

}
