package service

import (
	"github.com/jinzhu/gorm"
	"memo-api/model"
	"memo-api/pkg/utils"
	"memo-api/serializer"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	PassWord string `form:"password" json:"password" binding:"required,min=5,max=15"`
}

// Register 用户注册
func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{Status: 400, Msg: "用户已存在"}
	}
	user.UserName = service.UserName
	// 用户密码加密
	err := user.SetPassword(service.PassWord)
	if err != nil {
		return serializer.Response{
			Status: 400, Msg: err.Error(),
		}
	}
	// 创建用户
	err2 := model.DB.Create(&user).Error
	if err2 != nil {
		return serializer.Response{
			Status: 500, Msg: err2.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "用户注册成功",
	}
}

func (service *UserService) Login() serializer.Response {
	var user model.User
	// 判断用户是否存在
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "用户不存在",
			}
		}
		// 其他错误
		return serializer.Response{
			Status: 500,
			Msg:    err.Error(),
		}
	}
	if user.CheckPassword(service.PassWord) == false {
		return serializer.Response{Status: 400, Msg: "密码错误"}
	}
	// 生成用户token
	token, err := utils.CreateToken(user.ID, service.UserName, service.PassWord)
	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "Token签发错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
		Msg: "登录成功",
	}
}
