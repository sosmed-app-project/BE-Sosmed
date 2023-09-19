package levels

import "time"

type RoleCore struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RoleDataInterface interface {
	SelectAll(ID uint, Name string) ([]RoleCore, error)
	UpdateById(id uint, input RoleCore) (RoleCore, error)
}

type RoleServiceInterface interface {
	GetAll(ID uint, Name string) ([]RoleCore, error)
	UpdateRoleById(ID uint, input RoleCore) (RoleCore, error)
}
