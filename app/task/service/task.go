package service

import (
	"context"
	"encoding/json"
	"gin-gorm-memo/v2/app/task/repository/db/dao"
	"gin-gorm-memo/v2/app/task/repository/db/model"
	"gin-gorm-memo/v2/app/task/repository/db/mq"
	"gin-gorm-memo/v2/idl/pb"
	"gin-gorm-memo/v2/pkg/e"
	log "github.com/CocaineCong/micro-todoList/pkg/logger"
	"sync"
)

var TaskSrvIns *TaskSrv
var TaskSrvOnce sync.Once

type TaskSrv struct {
}

func GetTaskSrv() *TaskSrv {
	TaskSrvOnce.Do(func() {
		TaskSrvIns = &TaskSrv{}
	})
	return TaskSrvIns
}

func TaskMQ2DB(ctx context.Context, req *pb.TaskRequest) error {
	m := &model.Task{
		Uid:       uint(req.Uid),
		Title:     req.Title,
		Status:    int(req.Status),
		Content:   req.Content,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}
	return dao.NewTaskDao(ctx).CreateTask(m)
}

// CreateTask 创建备忘录，将备忘录信息生产，放到rabbitMQ消息队列中
func (t *TaskSrv) CreateTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	body, _ := json.Marshal(req) // title，content
	resp.Code = e.Success
	err = mq.SendMessage2MQ(body)
	if err != nil {
		resp.Code = e.Error
		return
	}
	return
}

// GetTaskList 实现备忘录服务接口 获取备忘录列表
func (t *TaskSrv) GetTaskList(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskListResponse) (err error) {
	resp.Code = e.Success
	if req.Limit == 0 {
		req.Limit = 10
	}
	// 查找备忘录
	r, count, err := dao.NewTaskDao(ctx).ListTaskByUserId(req.Uid, int(req.Start), int(req.Limit))
	if err != nil {
		resp.Code = e.Error
		log.LogrusObj.Error("ListTaskByUserId err:%v", err)
		return
	}
	// 返回proto里面定义的类型
	var taskRes []*pb.TaskModel
	for _, item := range r {
		taskRes = append(taskRes, BuildTask(item))
	}
	resp.TaskList = taskRes
	resp.Count = uint32(count)
	return
}

// GetTask 获取详细的备忘录
func (t *TaskSrv) GetTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code = e.Success
	r, err := dao.NewTaskDao(ctx).GetTaskByTaskIdAndUserId(req.Id, req.Uid)
	if err != nil {
		resp.Code = e.Error
		log.LogrusObj.Error("GetTask err:%v", err)
		return
	}
	taskRes := BuildTask(r)
	resp.TaskDetail = taskRes
	return
}

// UpdateTask 修改备忘录
func (t *TaskSrv) UpdateTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	// 查找该用户的这条信息
	resp.Code = e.Success
	err = dao.NewTaskDao(ctx).UpdateTask(req)
	if err != nil {
		resp.Code = e.Error
		log.LogrusObj.Error("UpdateTask err:%v", err)
		return
	}
	return
}

// DeleteTask 删除备忘录
func (t *TaskSrv) DeleteTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code = e.Success
	err = dao.NewTaskDao(ctx).DeleteTaskByTaskIdAndUserId(req.Id, req.Uid)
	if err != nil {
		resp.Code = e.Error
		log.LogrusObj.Error("DeleteTask err:%v", err)
		return
	}
	return
}

func BuildTask(item *model.Task) *pb.TaskModel {
	taskModel := pb.TaskModel{
		Id:          uint64(item.ID),
		Uid:         uint64(item.Uid),
		Title:       item.Title,
		Content:     item.Content,
		StartTime:   item.StartTime,
		EndTime:     item.EndTime,
		Status:      int64(item.Status),
		CreatedTime: item.CreatedAt.Unix(),
		UpdatedTime: item.UpdatedAt.Unix(),
	}
	return &taskModel
}
