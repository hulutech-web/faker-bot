package entities

import (
	"database/sql"
	"time"
)

type Category struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	PID       int64        `gorm:"column:pid" form:"pid" json:"pid"`
	Name      string       `gorm:"column:name" form:"name" json:"name"`
}
type Categories []Category
