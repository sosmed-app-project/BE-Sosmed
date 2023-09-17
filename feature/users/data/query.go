package data

import (
	"errors"
	"hris-app-golang/feature/users"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.UserDataInterface {
	return &userQuery{
		db: db,
	}
}

func (repo *userQuery) SelectById(id string) (users.UserCore, error) {
	var result User
	tx := repo.db.Preload("Role").Preload("Division").Preload("UserImport").Find(&result, id)
	if tx.Error != nil {
		return users.UserCore{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return users.UserCore{}, errors.New("data not found")
	}

	resultCore := ModelToCore(result)
	return resultCore, nil
}
