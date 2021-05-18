package main

import (
	"context"
	"fmt"
"github.com/coreos/etcd/clientv3"
"time"
)

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err   error
		kv clientv3.KV
		getResp *clientv3.GetResponse
	)

	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	if client, err = clientv3.New(config); err != nil{
		fmt.Println(err)
		return
	}
	kv = clientv3.NewKV(client)
	if getResp,err = kv.Get(context.TODO(), "/cron/jobs/job5", /*clientv3.WithCountOnly()*/); err != nil {
		fmt.Println(err)
		return
	}else{
		fmt.Println("Revision", getResp.Kvs, getResp.Count)
	}

	// 学习的方法，搜索博客，是为了快速的了解etcd v3 如何使用
}