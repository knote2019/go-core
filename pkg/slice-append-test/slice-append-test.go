package main

import "fmt"

func main() {
	slice1 := []string{"aaa"}
	fmt.Printf("value old addr = %p {%v}\n", &slice1, slice1)
	slice1 = append(slice1, "bbb")
	fmt.Printf("value new addr = %p {%v}\n", &slice1, slice1)

	slice2 := []string{"aaa"}
	fmt.Printf("value old addr = %p {%v}\n", slice2, slice2)
	slice2 = append(slice2, "bbb")
	fmt.Printf("value new addr = %p {%v}\n", slice2, slice2)
}

//value old addr = 0xc000004480 {[aaa]}
//value new addr = 0xc000004480 {[aaa bbb]}
//value old addr = 0xc00003c230 {[aaa]}
//value new addr = 0xc000004560 {[aaa bbb]}

// NOTE:
// append will create new slice.
// means same value(store pointer), but pointer changed.
