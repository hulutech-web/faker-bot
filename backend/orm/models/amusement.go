package models

import (
	"database/sql"
	"time"
)

type Amusement struct {
	ID        int64  `json:"id"`
	ItemCover string `gorm:"column:item_cover;" form:"item_cover" json:"item_cover"`
	ShareURL  string `gorm:"column:share_url" form:"share_url" json:"share_url" `
	Title     string `gorm:"title" form:"title" json:"title"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
