package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"memo-api/pkg/utils"
	"time"
)

// JWT token中间件验证
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200 // 状态码
		token := c.GetHeader("Authorization")
		fmt.Println("中间件token验证", token)
		if token == "" {
			code = 404
		} else {
			claims, err := utils.CheckToken(token)
			if err != nil {
				code = 403 // 无权限
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 401 // token过期
			}
		}
		if code != 200 {
			c.JSON(200, gin.H{"status": code, "msg": "token解析错误"})
			c.Abort()
			return
		}
		c.Next()
	}
}
