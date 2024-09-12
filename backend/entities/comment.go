package entities

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime `gorm:"index"`
	VideoID    uint         `gorm:"column:video_id" form:"video_id" json:"video_id"`
	UserName   string       `gorm:"column:user_name" form:"user_name" json:"user_name"`
	UserAvatar string       `gorm:"column:user_avatar" form:"user_avatar" json:"user_avatar"`
	Content    string       `json:"content"`
}
type Comments []Comment
