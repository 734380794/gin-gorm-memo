package rpc

import (
	"context"
	"gin-gorm-memo/v2/idl/pb"
)

func CreateTask(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	resp, err = TaskService.CreateTask(ctx, req)
	if err != nil {
		return
	}
	return
}

func UpdateTask(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	resp, err = TaskService.UpdateTask(ctx, req)
	if err != nil {
		return
	}
	return
}
func GetTaskList(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskListResponse, err error) {
	resp, err = TaskService.GetTaskList(ctx, req)
	if err != nil {
		return
	}
	return
}

func GetTask(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	resp, err = TaskService.GetTask(ctx, req)
	if err != nil {
		return
	}
	return
}

func DeleteTask(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	resp, err = TaskService.DeleteTask(ctx, req)
	if err != nil {
		return
	}
	return
}
