package http

import (
	"gin-gorm-memo/v2/app/gateway/rpc"
	"gin-gorm-memo/v2/idl/pb"
	"gin-gorm-memo/v2/pkg/ctl"
	"gin-gorm-memo/v2/pkg/utils"
	"gin-gorm-memo/v2/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegisterHandler(ctx *gin.Context) {
	var req pb.UserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserRegisterHandler-ShouldBind-rpc"))
		return
	}
	userResp, err := rpc.UserRegister(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserRegisterHandler-UserRegister-rpc"))
		return
	}
	ctx.JSON(http.StatusOK, userResp)
}

func UserLoginHandler(ctx *gin.Context) {
	var req pb.UserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserRegisterHandler-ShouldBind-rpc"))
		return
	}
	resp, err := rpc.UserLogin(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserRegisterHandler-UserLogin-rpc"))
		return
	}
	token, err := utils.GenerateToken(uint(resp.UserDetail.Id))
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserRegisterHandler-GenerateToken"))
		return
	}
	res := &types.TokenData{
		User:  resp,
		Token: token,
	}
	ctx.JSON(http.StatusOK, res)
}
