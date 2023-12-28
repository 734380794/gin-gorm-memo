package dao

import "gin-gorm-memo/v2/app/user/repository/db/model"

func migration() {
	// 自动迁移模式
	_db.Set("gorm:table_options", "charset=utf8").AutoMigrate(&model.User{})
}
