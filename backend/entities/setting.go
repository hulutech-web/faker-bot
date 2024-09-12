package entities

import (
	"database/sql"
	"time"
)

type Setting struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	Name      string       `json:"name"`
	Value     string       `json:"value"`
}
type Settings []*Setting
