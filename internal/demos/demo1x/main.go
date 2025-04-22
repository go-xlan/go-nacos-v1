package main

import (
	"context"
	"time"

	"github.com/go-xlan/go-nacos-v1/nacosv1"
	"github.com/yyle88/must"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
)

func main() {
	config := &nacosv1.Config{
		Endpoint:  "127.0.0.1:8848",
		AppName:   "demo1x",
		Address:   "0.0.0.0:8080",
		Group:     "DEFAULT_GROUP",
		Namespace: "public",
	}
	client := rese.P1(nacosv1.NewNacosClient(config, zaplog.ZAP))
	must.Done(client.RegisterService())
	client.Online(context.Background())

	time.Sleep(time.Minute)

	client.Offline(context.Background())
	must.Done(client.DeregisterService())
}
