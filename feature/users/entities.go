package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"not null"`
	TanggalLahir string `gorm:"not null"`
	Email        string `gorm:"not null"`
	NoHandphone  string `gorm:"not null"`
	Password     string `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
