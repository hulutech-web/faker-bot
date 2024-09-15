package models

import (
	"database/sql"
	"time"
)

// 上升词
type Trend struct {
	ID        uint   `gorm:"primarykey"`
	HotLevel  int    `gorm:"column:hot_level" json:"hot_level" form:"hot_level"`
	Sentence  string `json:"column:sentence" json:"sentence" form:"sentence"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
