package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	user := User{
		Name:    "kenny",
		Age:     30,
		Address: "shanghai",
	}

	reqBodyBytes := new(bytes.Buffer)
	err := json.NewEncoder(reqBodyBytes).Encode(user)
	if err != nil {
		panic(err)
	}
	url := "http://localhost:5555/users"
	req, err := http.NewRequest("POST", url, reqBodyBytes)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	fmt.Println("res status:", res.Status)
	fmt.Println("res headers:", res.Header)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("res body:", string(body))

}
