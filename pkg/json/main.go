package main

import (
	"fmt"
	"reflect"
)
func main() {
	type cat struct {
		Name string
		Type int `json:"type" id:"100"`
		//Type int `json: "type" id:"100"` （错误，不能有空格）
	}
	typeOfCat := reflect.TypeOf(cat{})
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		fmt.Println(catType.Tag.Get("json"))
	}
}
