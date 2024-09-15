package models

import (
	"database/sql"
	"time"
)

type HotSentence struct {
	ID        int64  `json:"id"`
	Word      string `gorm:"column:word" json:"word" form:"word"`
	Title     string `gorm:"column:title" json:"title" form:"title"`
	Cover     string `gorm:"cover" json:"cover" form:"cover"`
	ShareUrl  string `gorm:"column:share_url" json:"share_url" form:"share_url"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
