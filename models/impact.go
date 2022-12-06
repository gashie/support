package models

import "gorm.io/gorm"

type Impact struct {
	Title        string `json:"title"`
	Description string `json:"description"`
	gorm.Model
}
