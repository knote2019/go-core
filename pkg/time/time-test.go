package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("timestamp(secondes):%v;\n", time.Now().Unix())
	fmt.Printf("timestamp(secondes):%v;\n", time.Now().UnixNano()/1e9)
	fmt.Printf("timestamp(m secondes):%v;\n", time.Now().UnixNano()/1e6)
	fmt.Printf("timestamp(n secondes):%v;\n", time.Now().UnixNano())

	// sleep 5 seconds
	fmt.Println(time.Now())
	time.Sleep(5000000000)
	fmt.Println(time.Now())
	time.Sleep(5e9)
	fmt.Println(time.Now())
	time.Sleep(5 * time.Second)
	fmt.Println(time.Now())

	// 2006是go创建的时间，相当于特殊字符串。
	x := "2020-10-31 18:01:01"
	p, _ := time.Parse("2006-01-02 15:04:05", x)
	fmt.Println(p.Unix())

	t := time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05"))

	year1, month1, day1 := time.Now().Date()
	fmt.Printf("year: %v, type: %T \n", year1, year1)
	fmt.Printf("month: %v, type: %T \n", month1, month1)
	fmt.Printf("day: %v, type: %T \n", day1, day1)
}

// https://golang.google.cn/pkg/time/
