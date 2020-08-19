package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age int `min:"0" max:"100"`
}

func main()  {
	p := Person{Name: "kenny", Age: 30}
	pType := reflect.TypeOf(p)
	pValue := reflect.ValueOf(p)
	pNameField, ok := pType.FieldByName("Age")
	pNameFieldValue := pValue.FieldByName("Age")
	if ok {
		fmt.Println(pNameField.Tag.Get("min"))
		fmt.Println(pNameField.Tag.Get("max"))
		fmt.Println(pNameFieldValue.Int())
	}
}
