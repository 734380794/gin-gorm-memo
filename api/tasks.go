package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"memo-api/pkg/utils"
	"memo-api/service"
)

// CreateTask 创建任务
func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	token, _ := utils.CheckToken(c.GetHeader("Authorization"))

	if err := c.ShouldBind(&createTask); err == nil {
		res := createTask.Create(token.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

// ShowTask 查询任务
func ShowTask(c *gin.Context) {
	var showTask service.ShowTaskService

	if err := c.ShouldBind(&showTask); err == nil {
		res := showTask.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

// ListTask 任务列表
func ListTask(c *gin.Context) {
	var listTask service.ListTaskService
	token, _ := utils.CheckToken(c.GetHeader("Authorization"))

	if err := c.ShouldBind(&listTask); err == nil {
		res := listTask.List(token.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

// UpdateTask 更新任务
func UpdateTask(c *gin.Context) {
	var updateTask service.UpdateTaskService

	if err := c.ShouldBind(&updateTask); err == nil {
		res := updateTask.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

// DeleteTask 删除任务
func DeleteTask(c *gin.Context) {
	var deleteTask service.DeleteTaskService

	if err := c.ShouldBind(&deleteTask); err == nil {
		res := deleteTask.Delete(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}
