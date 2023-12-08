package posting

import (
	"time"

	"gorm.io/gorm"
)

type Posting struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	UserID    uint   `gorm:"not null"`
	Caption   string `gorm:"not null"`
	PhotoURL  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
