package task

import (
	"context"
	"encoding/json"
	"gin-gorm-memo/v2/app/task/repository/db/mq"
	"gin-gorm-memo/v2/app/task/service"
	"gin-gorm-memo/v2/consts"
	"gin-gorm-memo/v2/idl/pb"
)

type SyncTask struct {
}

func (s *SyncTask) RunTaskService(ctx context.Context) (err error) {
	rabbitMqQueue := consts.RabbitMqTaskQueue
	msgs, err := mq.ConsumeMessage(ctx, rabbitMqQueue)
	if err != nil {
		return
	}
	var forever chan struct{}
	go func() {
		for d := range msgs {
			req := new(pb.TaskRequest)
			err = json.Unmarshal(d.Body, req)
			if err != nil {
				return
			}
			err = service.TaskMQ2DB(ctx, req)
			if err != nil {
				return
			}
			d.Ack(false)
		}
	}()
	<-forever
	return nil
}
