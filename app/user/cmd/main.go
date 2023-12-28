package main

import (
	"fmt"
	"gin-gorm-memo/v2/app/user/repository/db/dao"
	"gin-gorm-memo/v2/config"
)

func main() {
	config.Init()
	dao.InitDB()
	fmt.Println("init")
}
