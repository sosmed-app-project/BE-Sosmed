package roles

import "time"

type RoleCore struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RoleDataInterface interface {
	GetAllRoles() ([]RoleCore, error)
}

type RoleServiceInterface interface {
	GetAllRoles() ([]RoleCore, error)
}
