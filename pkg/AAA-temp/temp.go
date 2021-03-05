package main

import (
	"fmt"
)

type Book struct {
	title   string
	author  string
	subject string
	bookId  int
}

func printBook(book *Book) {
	//fmt.Println(book.title)
	//fmt.Println(book.bookId)
}

func main() {
	b1 := Book{
		title:  "aaa",
		bookId: 1}

	printBook(&b1)
	slice1 := new([]int)
	slice2 := *new([]int)

	fmt.Println(slice1)
	fmt.Println(slice2)
	//person := make(map[string]string)
	//fmt.Println(*person) // invalid indirect of person
}
