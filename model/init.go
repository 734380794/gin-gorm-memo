package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 声明全局的gorm DB对象
var DB *gorm.DB

func Database(conn string) {
	// gorm 创建数据库连接
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err)
		fmt.Println("model/init.go 数据库连接错误")
	}
	fmt.Println("数据库连接成功")
	db.LogMode(true)
	// 如果是发型版本则不用输出
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	// 表名不加s
	db.SingularTable(true)
	// 设置db连接池
	db.DB().SetMaxIdleConns(20)
	// 设置db最大连接数
	db.DB().SetMaxOpenConns(100)
	// 最大连接时间 30
	db.DB().SetConnMaxLifetime(30)

	DB = db
	//migrate() 数据库迁移
}
