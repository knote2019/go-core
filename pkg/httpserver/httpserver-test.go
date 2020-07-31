package main

import (
	"fmt"
	"net/http"
)

//say hello to the world
func sayHello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello world"))
}

func main() {

	http.HandleFunc("/", sayHello)

	fmt.Println("begin")
	// 最后监听所有地址，方便在不同网络环境部署
	err := http.ListenAndServe("0.0.0.0:9000", nil)
	if err != nil {
		fmt.Printf("http.ListenAndServe() error: %v\n", err)
		return
	}
	fmt.Println("end")
}
