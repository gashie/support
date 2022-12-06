package models

import "gorm.io/gorm"

type Site struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	gorm.Model
}
