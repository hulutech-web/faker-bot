package models

import "gorm.io/gorm"

type Param struct {
	gorm.Model
	Url    string `gorm:"type:varchar(256);unique_index;column:url" form:"url" json:"url"`
	Method string `gorm:"type:varchar(16);column:method" form:"method" json:"method"`
	Body   string `gorm:"type:text;column:body" form:"body" json:"body"`
	Header string `gorm:"type:text;column:header" form:"header" json:"header"`
}
