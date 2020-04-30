package main

import (
	"fmt"
	"github.com/labstack/echo"
)

type subject struct {
	id  int
	str string
}

func main() {
	subj := subject{str: "world"}

	fmt.Println("hello %s", subj.id)
	fmt.Println("hello %s", subj.str)

	fmt.Println("hello $s", "kangminghong")

	var c1 Circle
	c1.radius = 12.0
	fmt.Println(c1.getArea())
}






