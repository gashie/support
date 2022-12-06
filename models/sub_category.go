package models

import "gorm.io/gorm"

type SubCategory struct {
	Title        string `json:"title"`
	Description string `json:"description"`
	gorm.Model
}
