package models

import (
	// "time"

	"time"

	// uuid "github.com/satori/go.uuid"
	// "gorm.io/gorm"
)

// type Base struct {
// 	ID        string     `gorm:"type:UUID;" json:"_id"`
// 	CreatedAt time.Time  `gorm:"type:Date;" json:"created_at"`
// 	UpdatedAt time.Time  `gorm:"type:Date;" json:"updated_at"`
// 	DeletedAt *time.Time `gorm:"type:Date;" json:"deleted_at"`
// }
type Base struct {
	// ID        uuid.UUID  `gorm:"type:char(36);primaryKey"`
	CreatedAt time.Time  `gorm:"type:timestamp;" json:"created_at"`
	UpdatedAt time.Time  `gorm:"type:Date;" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:Date;" json:"deleted_at"`
}
