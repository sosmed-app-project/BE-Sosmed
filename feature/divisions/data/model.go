package data

import (
	"time"

	"gorm.io/gorm"
)

type Division struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	gorm.DeletedAt
}
