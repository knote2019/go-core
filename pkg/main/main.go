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

	fmt.Println("hello", subj.id)
	fmt.Println("hello", subj.str)

	fmt.Println("hello", "kangminghong")

	// var c1 Circle
	// c1.radius = 12.0
	// fmt.Println(c1.getArea())
}
