package main

import "fmt"

func main()  {

	personMap := map[string]string{
		"France": "巴黎",
		"Italy": "罗马",
		"Japan": "东京",
	}
	fmt.Println(personMap["France"])

	funcMap := map[string][]string{
		"111": {"1111", "11111"},
		"222": {"2222", "22222"},
	}
	fmt.Println(funcMap)

}
