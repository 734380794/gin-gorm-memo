package script

import (
	"context"
	"gin-gorm-memo/v2/app/task/repository/db/mq/task"
)

func TaskCreateSync(ctx context.Context) {
	tSync := new(task.SyncTask)
	err := tSync.RunTaskService(ctx)
	if err != nil {
		return
	}
}
