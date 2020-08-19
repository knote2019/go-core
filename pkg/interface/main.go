package main

import "fmt"

type EI interface{}

type Adder interface {
	add() int
}

type MyStruct struct {
	X, Y int
}

func (a *MyStruct) add() int {
	return a.X + a.Y
}

func (s *MyStruct) String() string {
	return fmt.Sprintf("String [X: %v, Y: %v]", s.X, s.Y)
}

func (s *MyStruct) Error() string {
	return fmt.Sprintf("Error [X: %v, Y: %v]", s.X, s.Y)
}

func main() {
	// 声明一个空接口，可以任意赋值
	var i EI
	i = 42
	fmt.Printf("%v, %T\n", i, i)
	i = "hello"
	fmt.Printf("%v, %T\n", i, i)

	// 接口
	f := Adder(&MyStruct{3, 4})
	fmt.Println(f.add())

	// String(), Error()同时存在
	// 则会优先调用Error()
	fmt.Println(f)
}
