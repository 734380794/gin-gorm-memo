package rpc

import (
	"context"
	"gin-gorm-memo/v2/idl/pb"
)

func UserLogin(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	resp, err = UserService.UserLogin(ctx, req)
	if err != nil {
		return
	}
	return
}

func UserRegister(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	resp, err = UserService.UserRegister(ctx, req)
	if err != nil {
		return
	}
	return
}
