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
