package main

import (
	"fmt"
	"github.com/masterzen/winrm"
	"os"
)

func main() {
	host := "127.0.0.1"
	port := 5985
	username := "emigkag"
	password := "6YHN%tgb"
	cmd := "ipconfig"
	endpoint := winrm.NewEndpoint(host, port, false, false, nil, nil, nil, 0)
	client, err := winrm.NewClient(endpoint, username, password)
	if err != nil {
		panic(err)
	}
	aaa, bbb := client.Run(cmd, os.Stdout, os.Stderr)
	fmt.Println(aaa)
	fmt.Println(bbb)
}
