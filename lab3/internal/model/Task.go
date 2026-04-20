package model

import (
	"time"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;not null"`
	Description string
	Date        time.Time
}