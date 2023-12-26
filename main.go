package main

import (
	"gin-gorm-memo/conf"
	"gin-gorm-memo/routes"
)

func main() {
	conf.Init()

	r := routes.NewRouter()
	r.Run(conf.HttpPort)
}
