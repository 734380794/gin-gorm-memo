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
	Status  int    `json:"status" form:"status"`
}

type UpdateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"`
}

type ShowTaskService struct {
}

type DeleteTaskService struct {
}

type ListTaskService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
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

// Show 显示任务
func (service *ShowTaskService) Show(tid string) serializer.Response {
	var task model.Task
	code := 200
	err := model.DB.First(&task, tid).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "数据查询失败",
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
	}
}

// List 任务列表
func (service *ListTaskService) List(uid uint) serializer.Response {
	count := 0
	var tasks []model.Task
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Count(&count).Limit(service.PageSize).Offset(service.PageNum - 1).Find(&tasks)
	return serializer.BuildListResponse(tasks, uint(count))
}

// Update 更新任务
func (service *UpdateTaskService) Update(tid string) serializer.Response {
	var task model.Task
	model.DB.First(&task, tid)
	task.Title = service.Title
	task.Content = service.Content
	task.Status = service.Status
	model.DB.Save(task)
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildTask(task),
		Msg:    "更新成功",
	}
}

// Delete 删除任务
func (service *DeleteTaskService) Delete(id string) serializer.Response {
	var task model.Task
	err := model.DB.Delete(&task, id).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "删除失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "删除成功",
	}
}
