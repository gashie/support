package models

import "gorm.io/gorm"

type Category struct {
	Title        string `json:"title"`
	Description string `json:"description"`
	gorm.Model
}
