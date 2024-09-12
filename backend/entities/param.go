package entities

import (
	"database/sql"
	"time"
)

type Param struct {
	ID        uint         `gorm:"primarykey"`
	CreatedAt time.Time    `gorm:"column:created_at" form:"created_at" json:"created_at"`
	UpdatedAt time.Time    `gorm:"column:updated_at" form:"updated_at" json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at" form:"deleted_at" json:"deleted_at"`
	Url       string       `gorm:"type:varchar(256);unique_index;column:url" form:"url" json:"url"`
	Method    string       `gorm:"type:varchar(16);column:method" form:"method" json:"method"`
	Body      string       `gorm:"type:text;column:body" form:"body" json:"body"`
	Header    string       `gorm:"type:text;column:header" form:"header" json:"header"`
	Label     string       `gorm:"type:varchar(64);column:label" form:"label" json:"label"`
}
