package serializer

import "memo-api/model"

type User struct {
	ID       uint   `json:"id" form:"id" example:"1"`                   // 用户ID
	UserName string `json:"user_name" form:"user_name" example:"kevin"` // 用户名
	Status   string `json:"status" form:"status" example:"0"`           // 用户状态
	CreateAt int64  `json:"create_at" form:"create_at"`                 // 用户状态
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
	}
}
