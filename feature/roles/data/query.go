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

func (repo *RoleQuery) SelectAll(ID uint, Name string) ([]levels.RoleCore, error) {

	var roleData []Role
	var tx *gorm.DB

	if ID != 0 {
		tx = repo.db.Where("id = ?", ID).Find(&roleData)
	} else {
		tx = repo.db.Find(&roleData)
	}
	if tx.Error != nil {
		return nil, tx.Error
	}

	if Name != "" {
		tx = repo.db.Where("name = ?", Name).Find(&roleData)
	} else {
		tx = repo.db.Find(&roleData)
	}
	if tx.Error != nil {
		return nil, tx.Error
	}

	var rolesCore []levels.RoleCore
	for _, value := range roleData {
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
