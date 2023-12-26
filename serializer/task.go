package serializer

import "gin-gorm-memo/model"

type Task struct {
	ID        uint   `json:"id"` // 任务ID
	Title     string `json:"title"`
	Content   string `json:"content"`
	View      uint64 `json:"view"`
	Status    int    `json:"status"`
	CreateAt  int64  `json:"created_at"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

// BuildTask 序列化用户
func BuildTask(item model.Task) Task {
	return Task{
		ID:        item.ID,
		Title:     item.Title,
		Content:   item.Content,
		Status:    item.Status,
		CreateAt:  item.CreateAt,
		StartTime: item.StartTime,
		EndTime:   item.EndTime,
	}
}
