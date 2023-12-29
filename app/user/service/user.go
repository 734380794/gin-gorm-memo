package service

import (
	"context"
	"errors"
	"gin-gorm-memo/v2/app/user/repository/db/dao"
	"gin-gorm-memo/v2/app/user/repository/db/model"
	"gin-gorm-memo/v2/idl/pb"
	"gin-gorm-memo/v2/pkg/e"
	"gorm.io/gorm"
	"sync"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct {
}

// GetUserSrv 懒汉式的单例模式
func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

// GetUserSrvHuqury 饿汉式的单例模式
func GetUserSrvHuqury() *UserSrv {
	if UserSrvIns != nil {
		return new(UserSrv)
	}
	return UserSrvIns
}

func (u *UserSrv) UserLogin(ctx context.Context, req *pb.UserRequest, resp *pb.UserResponse) (err error) {
	resp.Code = e.Success
	user, err := dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if err != nil {
		resp.Code = e.Error
		return
	}

	if !user.CheckPassword(req.Password) {
		resp.Code = e.InvalidParams
		return
	}

	resp.UserDetail = BuildUser(user)
	return
}

func (u *UserSrv) UserRegister(ctx context.Context, req *pb.UserRequest, resp *pb.UserResponse) (err error) {
	if req.Password != req.PasswordConfirm {
		err = errors.New("两次密码输入不一致")
		return
	}
	resp.Code = e.Success
	_, err = dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if err != nil {
		if err == gorm.ErrRecordNotFound { // 如果不存在就继续下去
			// ...continue
		} else {
			resp.Code = e.Error
			return
		}
	}
	user := &model.User{
		UserName: req.UserName,
	}
	// 加密密码
	if err = user.SetPassword(req.Password); err != nil {
		resp.Code = e.Error
		return
	}
	if err = dao.NewUserDao(ctx).CreateUser(user); err != nil {
		resp.Code = e.Error
		return
	}

	resp.UserDetail = BuildUser(user)
	return
}

func BuildUser(item *model.User) *pb.UserModel {
	userModel := pb.UserModel{
		Id:        uint32(item.ID),
		UserName:  item.UserName,
		CreatedAt: string(item.CreatedAt.Unix()),
		UpdatedAt: string(item.UpdatedAt.Unix()),
	}
	return &userModel
}
