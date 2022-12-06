package models

import "gorm.io/gorm"

type Status struct {
	Title        string `json:"title"`
	Description string `json:"description"`
	gorm.Model
}
