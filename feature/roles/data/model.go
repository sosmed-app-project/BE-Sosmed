package data

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	gorm.DeletedAt
}
