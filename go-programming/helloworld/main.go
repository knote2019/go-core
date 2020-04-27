package main

import (
	"fmt"
)

type subject struct {
	id  int
	str string
}

func main() {
	subj := subject{str: "world"}

	fmt.Printf("hello %s", subj.id)
	fmt.Printf("hello %s", subj.str)

	fmt.Printf("hello $s", "kangminghong")
}


