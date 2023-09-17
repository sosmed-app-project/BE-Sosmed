package data

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        string         `gorm:"column:id;type:varchar(30);primaryKey"`
	Name      string         `gorm:"column:name;not null"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}
