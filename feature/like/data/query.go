// like/data.go

package like

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type LikeDataInterface interface {
	GetLikes() ([]Like, error)
	CreateLike(input Like) error
	// Fungsi-fungsi lainnya sesuai kebutuhan
}

type LikeData struct {
	db *gorm.DB
}

func NewLikeData(db *gorm.DB) LikeDataInterface {
	return &LikeData{
		db: db,
	}
}

func (data *LikeData) GetLikes() ([]Like, error) {
	var likes []Like
	if err := data.db.Find(&likes).Error; err != nil {
		return nil, err
	}
	return likes, nil
}

func (data *LikeData) CreateLike(input Like) error {
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	tx := data.db.Create(&input)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}
	return nil
}

// Implementasikan fungsi-fungsi lainnya sesuai kebutuhan
