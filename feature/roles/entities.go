package levels

import "time"

type RoleCore struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RoleDataInterface interface {
	SelectAll() ([]RoleCore, error)
	UpdateById(id uint, input RoleCore) (RoleCore, error)
}

type RoleServiceInterface interface {
	GetAll() ([]RoleCore, error)
	UpdateRoleById(ID uint, input RoleCore) (RoleCore, error)
}
