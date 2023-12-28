package model

// 定义数据库model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	Password string
}
