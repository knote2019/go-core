package main

import (
	"encoding/json"
	"fmt"
)
func main() {
	m := map[string]string{"type":"10", "msg":"hello."}
	mjson, _ := json.Marshal(m)
	mString := string(mjson)
	fmt.Printf("print mString:%s",mString)
}
