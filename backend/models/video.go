package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Title string `gorm:"column:title" form:"title" json:"title"`
	Url   string `gorm:"url" form:"url" json:"url" `
	Desc  string `gorm:"desc" form:"desc" json:"desc"`
}
type VideoList []Video
