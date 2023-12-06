package data

import (
	levels "hris-app-golang/feature/roles"
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint           `gorm:"column:id;primaryKey"`
	Name      string         `gorm:"column:name;not null"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func RoleModelToCore(input Role) levels.RoleCore {
	var roleCore = levels.RoleCore{
		ID:        input.ID,
		Name:      input.Name,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}
	return roleCore
}

func RoleCoreToModel(dataCore levels.RoleCore) Role {
	return Role{
		ID:   dataCore.ID,
		Name: dataCore.Name,
	}
}
