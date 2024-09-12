package models

import "gorm.io/gorm"

type Setting struct {
	gorm.Model
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Settings []*Setting
