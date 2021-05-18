package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"time"
)

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err   error
		kv clientv3.KV
		delResp *clientv3.DeleteResponse
		kvpair *mvccpb.KeyValue
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
	if delResp,err = kv.Delete(context.TODO(), "/cron/jobs/job5", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
		return
	}
	if len(delResp.PrevKvs) != 0 {
		for _, kvpair = range delResp.PrevKvs {
			fmt.Println("删除了", string(kvpair.Key), string(kvpair.Value))
		}
	}

	// 学习的方法，搜索博客，是为了快速的了解etcd v3 如何使用
}