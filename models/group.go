package models

import "gorm.io/gorm"

type Group struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	gorm.Model
}
