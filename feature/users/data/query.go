package data

import (
	"errors"
	"hris-app-golang/feature/users"

	"gorm.io/gorm"
)

type userQuery struct {
	db        *gorm.DB
	dataLogin users.UserCore
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

func (repo *userQuery) DeleteById(id string) error {
	var userGorm User
	tx := repo.db.Where("id = ?", id).Delete(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *userQuery) Login(email string, password string) (dataLogin users.UserCore, err error) {

	var data User
	tx := repo.db.Where("email = ? and password = ?", email, password).Find(&data)
	if tx.Error != nil {
		return users.UserCore{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return users.UserCore{}, errors.New("data not found")
	}
	dataLogin = ModelToCore(data)
	repo.dataLogin = dataLogin
	return dataLogin, nil
}
