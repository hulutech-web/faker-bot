package entities

import (
	"database/sql"
	"time"
)

type Video struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	Title     string       `gorm:"column:title" form:"title" json:"title"`
	Url       string       `gorm:"url" form:"url" json:"url" `
	Desc      string       `gorm:"desc" form:"desc" json:"desc"`
}
type VideoList []Video
