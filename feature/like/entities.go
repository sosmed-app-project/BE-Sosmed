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
	Read() ([]DivisionCore, error)
}

type DivisionServiceInterface interface {
	GetDiv() ([]DivisionCore, error)
}
