package api

import (
	"github.com/gin-gonic/gin"
	"memo-api/service"
)

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	// 声明user服务对象
	var userRegister service.UserService
	// 绑定服务对象
	if err := c.ShouldBind(&userRegister); err == nil {
		// 执行注册方法
		res := userRegister.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
	// 声明user服务对象
	var userLogin service.UserService
	// 绑定服务对象
	if err := c.ShouldBind(&userLogin); err == nil {
		// 执行注册方法
		res := userLogin.Login()
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
