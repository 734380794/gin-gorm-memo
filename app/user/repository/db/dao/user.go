package dao

import (
	"context"
	"gin-gorm-memo/v2/app/user/repository/db/model"
	"gorm.io/gorm"
)

// 定义对数据库user表model的curd
type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClient(ctx)}
}

func (dao *UserDao) FindUserByUserName(userName string) (r *model.User, err error) {
	err = dao.Model(&model.User{}).Where("user_name=?", userName).Find(&r).Error
	//record not found
	if err != nil {
		return
	}
	return
}

func (dao *UserDao) CreateUser(in *model.User) (err error) {
	return dao.Model(&model.User{}).Create(&in).Error
}
