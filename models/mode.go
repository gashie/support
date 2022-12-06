package models

import "gorm.io/gorm"

type Mode struct {
	Title        string `json:"title"`
	Description string `json:"description"`
	gorm.Model
}
