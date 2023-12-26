package main

import (
	"memo-api/conf"
	"memo-api/routes"
)

func main() {
	conf.Init()

	r := routes.NewRouter()
	r.Run(conf.HttpPort)
}
