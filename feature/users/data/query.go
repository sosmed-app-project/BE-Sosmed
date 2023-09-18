package data

import (
	"errors"
	"hris-app-golang/feature/users"
	"hris-app-golang/helper"

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

// Insert implements users.UserDataInterface.
func (repo *UserQuery) Insert(input users.UserCore) error {
	var userModel = UserCoreToModel(input)
	var errGen error
	if errGen != nil {
		return errGen
	}
	hass, errHass := helper.HassPassword(userModel.Password)
	if errHass != nil {
		return errHass
	}
	userModel.Password = hass
	tx := repo.db.Create(&userModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}
	return nil
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

// SelectAll implements users.UserDataInterface.
func (repo *UserQuery) SelectAll(role_id string, division_id string) ([]users.UserCore, error) {
	var userModel []User
	if role_id == "3" {
		tx := repo.db.Where("role_id in (1,2)").Find(&userModel)
		if tx.Error != nil {
			return nil, tx.Error
		}
		if tx.RowsAffected == 0 {
			return nil, errors.New("no row affected")
		}
	} else if role_id == "1" || role_id == "2" {
		tx := repo.db.Where("role_id in (1,2) and division_id = ?", division_id).Find(&userModel)
		if tx.Error != nil {
			return nil, tx.Error
		}
		if tx.RowsAffected == 0 {
			return nil, errors.New("no row affected")
		}
	}
	var userCore []users.UserCore
	for _, value := range userModel {
		var user = ModelToCore(value)
		userCore = append(userCore, user)
	}
	return userCore, nil
}

// Update implements users.UserDataInterface.
func (repo *UserQuery) Update(id string, input users.UserCore) error {
	var userModel = UserCoreToModel(input)
	tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(userModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
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
