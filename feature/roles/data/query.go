package data

import (
	levels "hris-app-golang/feature/roles"

	"gorm.io/gorm"
)

type RoleQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) levels.RoleDataInterface {
	return &RoleQuery{
		db: db,
	}
}

func (repo *RoleQuery) SelectAll() ([]levels.RoleCore, error) {

	var rolesData []Role
	tx := repo.db.Find(&rolesData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var rolesCore []levels.RoleCore
	for _, value := range rolesData {
		rolesCore = append(rolesCore, levels.RoleCore{
			ID:   value.ID,
			Name: value.Name,
		})
	}
	return rolesCore, nil
}

func (repo *RoleQuery) UpdateById(ID uint, input levels.RoleCore) (levels.RoleCore, error) {

	roleGorm := RoleCoreToModel(input)
	tx := repo.db.Model(&Role{}).Where("id = ?", ID).Updates(roleGorm)
	if tx.Error != nil {
		return levels.RoleCore{}, tx.Error
	}
	return RoleModelToCore(roleGorm), nil
}
