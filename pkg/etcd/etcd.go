package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var(
	etcdConfig clientv3.Config
	etcdClient *clientv3.Client
	err error
)

func main() {
	etcdConfig = clientv3.Config{
		Endpoints: []string{"192.168.99.100:2379"},
		DialTimeout: 5 * time.Second,
		Username: "etcd",
		Password: "etcd1234",
	}
	etcdClient, err = clientv3.New(etcdConfig)
	if err != nil {
		fmt.Println(err)
	}

	//etcdClient.Sync()

}
