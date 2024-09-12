package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	VideoID    uint   `gorm:"column:video_id" form:"video_id" json:"video_id"`
	UserName   string `gorm:"column:user_name" form:"user_name" json:"user_name"`
	UserAvatar string `gorm:"column:user_avatar" form:"user_avatar" json:"user_avatar"`
	Content    string `json:"content"`
}
type Comments []Comment
