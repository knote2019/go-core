package main

import "fmt"

type Person struct {
	name   string
	age    int
	gender string
}

func main() {

	var p1 = new(Person)
	p1.name = "kenny"
	p1.age = 30
	p1.gender = "men"

	// p1 是一个指针
	fmt.Printf("%T\n", p1)
	fmt.Printf("%p\n", p1)  // p1 保存的地址
	fmt.Printf("%p\n", &p1) // p1 自己的地址

	var p2 = Person{
		name:   "kangminghong",
		age:    32,
		gender: "men",
	}

	// p1 是一个变量
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", &p2)
}
