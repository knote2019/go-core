package main

import "fmt"

func changeValue(v string) {
	fmt.Printf("value new addr = %p {%v}\n", &v, v)
	v = "111"
}

func changeArray(a [3]string) {
	fmt.Printf("array new addr = %p {%v}\n", &a, a)
	a[0] = "111"
}

func changeSlice(s []string) {
	fmt.Printf("slice new addr = %p {%v}\n", s, s)
	s[0] = "111"
}

func changeMap(m map[int]string) {
	fmt.Printf("map new addr = %p {%v}\n", m, m)
	m[0] = "111"
}

func main() {
	// transfer value for value and array
	value1 := "aaa"
	array1 := [3]string{"aaa", "bbb", "ccc"}

	// transfer pointer for slice and map
	slice1 := []string{"aaa", "bbb", "ccc"}
	map1 := map[int]string{
		0: "aaa",
		1: "bbb",
		2: "ccc"}

	fmt.Printf("value old addr = %p {%v}\n", &value1, value1)
	changeValue(value1)

	fmt.Printf("array old addr = %p {%v}\n", &array1, array1)
	changeArray(array1)

	fmt.Printf("slice old addr = %p {%v}\n", slice1, slice1)
	changeSlice(slice1)

	fmt.Printf("map old addr = %p {%v}\n", map1, map1)
	changeMap(map1)
}

//value old addr = 0xc00003c1f0 {aaa}
//value new addr = 0xc00003c210 {aaa}
//array old addr = 0xc00006e330 {[aaa bbb ccc]}
//array new addr = 0xc00006e3f0 {[aaa bbb ccc]}
//slice old addr = 0xc00006e360 {[aaa bbb ccc]}
//slice new addr = 0xc00006e360 {[aaa bbb ccc]}
//map old addr = 0xc00006e390 {map[0:aaa 1:bbb 2:ccc]}
//map new addr = 0xc00006e390 {map[0:aaa 1:bbb 2:ccc]}
