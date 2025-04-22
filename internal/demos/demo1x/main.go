package main

import (
	"context"
	"time"

	"github.com/go-xlan/go-nacos-v1/nacosv1"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
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
	clientOptions := []constant.ClientOption{
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogDir("/tmp/nacos/log"),
	}
	client := rese.P1(nacosv1.NewNacosClient(config, clientOptions, zaplog.ZAP.NewZap("module", "demo1x")))
	must.Done(client.RegisterService())
	client.Online(context.Background())

	time.Sleep(time.Minute)

	client.Offline(context.Background())
	must.Done(client.DeregisterService())
}
