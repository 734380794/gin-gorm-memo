package main

import (
	"fmt"
	"gin-gorm-memo/v2/app/user/repository/db/dao"
	"gin-gorm-memo/v2/app/user/service"
	"gin-gorm-memo/v2/config"
	"gin-gorm-memo/v2/idl/pb"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func main() {
	config.Init()
	dao.InitDB()
	// etcd注册
	etcdReg := registry.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)),
	)

	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcUserService"), // 微服务名字
		micro.Address(config.UserServiceAddress),
		micro.Registry(etcdReg), // etcd注册件
	)
	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = pb.RegisterUserServiceHandler(microService.Server(), service.GetUserSrv())
	// 启动微服务
	_ = microService.Run()
}
