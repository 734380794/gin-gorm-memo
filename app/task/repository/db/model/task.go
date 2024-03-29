package model

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"index;not null"`
	Status    int    `gorm:"default:0"`
	Content   string `gorm:"text"`
	StartTime int64
	EndTime   int64
}
