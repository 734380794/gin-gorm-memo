package dao

import (
	"context"
	"gin-gorm-memo/v2/app/task/repository/db/model"
	"gin-gorm-memo/v2/idl/pb"
	"gorm.io/gorm"
)

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(ctx context.Context) *TaskDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &TaskDao{NewDBClient(ctx)}
}

func (dao *TaskDao) CreateTask(in *model.Task) error {
	return dao.Model(&model.Task{}).Create(in).Error
}

func (dao *TaskDao) ListTaskByUserId(userId uint64, start, limit int) (r []*model.Task, count int64, err error) {
	err = dao.Model(&model.Task{}).Offset(start).Limit(limit).Where("uid=?", userId).Find(&r).Error
	if err != nil {
		return
	}
	err = dao.Model(&model.Task{}).Where("uid=?", userId).Count(&count).Error
	return
}

func (dao *TaskDao) GetTaskByTaskIdAndUserId(Tid, Uid uint64) (r *model.Task, err error) {
	err = dao.Model(&model.Task{}).Where("id = ? AND uid =?", Tid, Uid).Find(&r).Error
	return
}
func (dao *TaskDao) UpdateTask(req *pb.TaskRequest) (err error) {
	var r *model.Task
	err = dao.Model(&model.Task{}).Where("id = ? AND uid =?", req.Id, req.Uid).Find(&r).Error
	if err != nil {
		r.Title = req.Title
		r.Content = req.Content
		r.Status = int(req.Status)
	}
	return dao.Save(&r).Error
}
func (dao *TaskDao) DeleteTaskByTaskIdAndUserId(Tid, Uid uint64) (err error) {
	err = dao.Model(&model.Task{}).Where("id = ? AND uid =?", Tid, Uid).Delete(&model.Task{}).Error
	return
}
