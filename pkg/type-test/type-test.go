package main

import (
	"fmt"
)

func MyPrintf(args ...interface{}) {
	for index, arg := range args {
		fmt.Println(index)
		switch v := arg.(type) {
		case string:
			fmt.Println(arg, "is string value.")
			fmt.Println(v)
		case int:
			fmt.Println(arg, "is int value.")
			fmt.Println(v)
		default:
			fmt.Println(arg, "is unknown type.")
			fmt.Println(v)
		}
	}
}

func main() {
	MyPrintf("kenny", 30)
}
