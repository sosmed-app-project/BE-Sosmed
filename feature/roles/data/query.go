package data

import (
	"hris-app-golang/feature/roles"

	"gorm.io/gorm"
)

type RoleQuery struct {
	db *gorm.DB
}

// GetAllRoles implements roles.RoleDataInterface.
func (res *RoleQuery) GetAllRoles() ([]roles.RoleCore, error) {
	var roleData []Role
	tx := res.db.Find(roleData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var rolesCoreData []roles.RoleCore
	for _, value := range roleData {
		roleCore := roles.RoleCore{
			ID:        value.ID,
			Name:      value.Name,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		rolesCoreData = append(rolesCoreData, roleCore)
	}
	return rolesCoreData, nil
}

func New(db *gorm.DB) roles.RoleDataInterface {
	return &RoleQuery{
		db: db,
	}
}
