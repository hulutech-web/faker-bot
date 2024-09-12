package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	PID  int64  `gorm:"column:pid" form:"pid" json:"pid"`
	Name string `gorm:"column:name" form:"name" json:"name"`
}
type Categories []Category
