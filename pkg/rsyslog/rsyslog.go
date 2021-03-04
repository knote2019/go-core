package main

import (
	"log"
	"log/syslog"
)

// TODO: fix bug
func main() {
	// 如果network是空，Dial会连接到本地的syslog服务器。
	sysLog, err := syslog.Dial("", "", syslog.LOG_ERR, "Saturday")
	if err != nil {
		log.Fatal(err)
	}
	sysLog.Emerg("Hello world!")
}
