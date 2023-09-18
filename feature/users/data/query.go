package data

import (
	"errors"
	"hris-app-golang/feature/users"
	"hris-app-golang/helper"

	"gorm.io/gorm"
)

type UserQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.UserDataInterface {
	return &UserQuery{
		db: db,
	}
}

// Delete implements users.UserDataInterface.
func (*UserQuery) Delete(id string) error {
	panic("unimplemented")
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

// Login implements users.UserDataInterface.
func (repo *UserQuery) Login(email string, password string) (users.UserCore, error) {
	var userData users.UserCore
	var err error
	return userData, err
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

// SelectById implements users.UserDataInterface.
func (*UserQuery) SelectById(id string) (users.UserCore, error) {
	panic("unimplemented")
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
