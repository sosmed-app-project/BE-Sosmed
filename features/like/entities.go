package like

import (
	"time"

	"gorm.io/gorm"
)

type Like struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	UserID    uint `gorm:"not null"`
	PostingID uint `gorm:"not null"`
	CountLike bool `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
