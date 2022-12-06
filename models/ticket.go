package models

import "gorm.io/gorm"

type Ticket struct {
	Subject        string `json:"subject"`
	Description string `json:"description"`
	gorm.Model
}
