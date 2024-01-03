package http

import (
	"gin-gorm-memo/v2/app/gateway/rpc"
	"gin-gorm-memo/v2/idl/pb"
	"gin-gorm-memo/v2/pkg/ctl"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-ShouldBind"))
		return
	}
	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-GetUserInfo"))
		return
	}
	req.Uid = uint64(user.Id)
	resp, err := rpc.CreateTask(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-CreateTask"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
}

func UpdateTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UpdateTaskHandler-ShouldBind"))
		return
	}
	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UpdateTaskHandler-GetUserInfo"))
		return
	}
	req.Uid = uint64(user.Id)
	resp, err := rpc.UpdateTask(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UpdateTaskHandler-UpdateTask"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
}
func GetTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "GetTaskHandler-ShouldBind"))
		return
	}
	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "GetTaskHandler-GetUserInfo"))
		return
	}
	req.Uid = uint64(user.Id)
	resp, err := rpc.GetTask(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "GetTaskHandler-GetTask"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
}
func GetTaskListHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "GetTaskListHandler-ShouldBind"))
		return
	}
	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "GetTaskListHandler-GetUserInfo"))
		return
	}
	req.Uid = uint64(user.Id)
	resp, err := rpc.GetTaskList(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "GetTaskListHandler-GetTaskList"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
}
func DeleteTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "DeleteTaskHandler-ShouldBind"))
		return
	}
	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "DeleteTaskHandler-GetUserInfo"))
		return
	}
	req.Uid = uint64(user.Id)
	resp, err := rpc.DeleteTask(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "DeleteTaskHandler-DeleteTask"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
}
