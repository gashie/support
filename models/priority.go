package models

import "gorm.io/gorm"

type Priority struct {
	Title        string `json:"title"`
	Description string `json:"description"`
	gorm.Model
}
