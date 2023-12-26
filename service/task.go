package service

import (
	"memo-api/model"
	"memo-api/serializer"
	"time"
)

// CreateTaskService task结构体
type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  string `json:"status" form:"status"`
}

// Create 创建任务
func (service *CreateTaskService) Create(id uint) serializer.Response {
	var code = 200
	var user model.User
	model.DB.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Status:    0,
		Content:   service.Content,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "添加数据失败",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "添加数据成功",
	}

}
