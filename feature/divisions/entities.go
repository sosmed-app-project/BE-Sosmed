package divisions

import (
	"time"
)

type DivisionCore struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DivisionDataInterface interface {
    Insert(division DivisionCore) error
    SelectAll() ([]DivisionCore, error)
    SelectByID(id uint) (DivisionCore, error)
    Update(id uint, division DivisionCore) error
    Delete(id uint) error
}