package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"memo-api/api"
	"memo-api/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))

	v1 := r.Group("api/v1")
	{
		// 用户注册
		v1.POST("user/register", api.UserRegister)
		// 用户登录
		v1.POST("user/login", api.UserLogin)
		authed := v1.Group("/")
		// token 验证中间件
		authed.Use(middleware.JWT())
		{
			// 创建任务
			v1.POST("task/create", api.CreateTask)
			// 查询任务
			v1.GET("task/show/:id", api.ShowTask)
		}
	}
	return r
}
