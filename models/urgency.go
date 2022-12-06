package models

import "gorm.io/gorm"

type Urgency struct {
	Title        string `json:"title"`
	Description string `json:"description"`
	gorm.Model
}
