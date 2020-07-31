package main

import "fmt"

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func main() {

	const (
		C0 = iota
		C1 = iota
		C2 = iota
	)

	// 可以简化为下面的写法：
	// const (
	// 	C0 = iota
	// 	C1
	// 	C2
	// )

	fmt.Println(C0, C1, C2) // "0 1 2"

	const (
		B1 = iota + 1
		_
		B3
		B4
	)
	fmt.Println(B1, B3, B4)

	var d Direction = North
	fmt.Print(d)
	switch d {
	case North:
		fmt.Println(" goes up.")
	case South:
		fmt.Println(" goes down.")
	default:
		fmt.Println(" stays put.")
	}

}
