package divisions

import "time"

type DivisionCore struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
