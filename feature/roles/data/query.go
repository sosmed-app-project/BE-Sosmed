package data

import (
	"hris-app-golang/feature/roles"

	"gorm.io/gorm"
)

type RoleQuery struct {
	db *gorm.DB
}

func NewRoleQuery(db *gorm.DB) roles.RoleDataInterface {
	return &RoleQuery{
		db: db,
	}
}

func (repo *RoleQuery) Read() ([]roles.RoleCore, error) {
	var RoleData []Role
	tx := repo.db.Find(&RoleData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var roleCore []roles.RoleCore
	for _, value := range RoleData {
		roleCore = append(roleCore, roles.RoleCore{
			ID:   value.ID,
			Name: value.Name,
		})
	}
	return roleCore, nil
}
