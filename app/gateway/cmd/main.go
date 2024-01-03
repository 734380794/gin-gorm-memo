package main

import (
	"fmt"
	"gin-gorm-memo/v2/app/gateway/router"
	"gin-gorm-memo/v2/app/gateway/rpc"
	"gin-gorm-memo/v2/config"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
	"time"
)

func main() {
	// 配置文件初始化
	config.Init()
	// 网关注册
	rpc.InitRPC()
	// etcd注册
	etcdReg := registry.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)),
	)

	// 得到一个微服务实例
	webService := web.NewService(
		web.Name("httpService"), // 微服务名字
		web.Address("localhost:4000"),
		web.Registry(etcdReg), // etcd注册件
		web.Handler(router.NewRouter()),
		web.RegisterTTL(time.Second*30),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	// 结构命令行参数，初始化
	webService.Init()
	// 启动微服务
	_ = webService.Run()
}
