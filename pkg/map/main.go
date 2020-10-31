package main

import (
	"fmt"
	"reflect"
)

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


	m1:=map[string]interface{}{"a":"1", "b":2, "c":3};
	m2:=map[string]interface{}{"a":1, "c":"3", "b":2};

	fmt.Println(reflect.DeepEqual(m1["a"],m2["a"]))
	fmt.Println(reflect.DeepEqual(m1["b"],m2["b"]))
}
