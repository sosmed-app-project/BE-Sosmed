package data

import (
	"time"

	"gorm.io/gorm"
)

type Division struct {
	ID        uint           `gorm:"column:id;primaryKey"`
	Name      string         `gorm:"column:name;not null"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}
