package data

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID          string
	Name        string
	Address     string
	PhoneNumber string
	Email       string
	Website     string
	Founder     string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	gorm.DeletedAt
}
