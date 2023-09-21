package data

import (
	"errors"
	"fmt"
	"hris-app-golang/feature/users"
	"hris-app-golang/helper"
	"reflect"

	"gorm.io/gorm"
)

type UserQuery struct {
	db        *gorm.DB
	dataLogin users.UserCore
}

func New(db *gorm.DB) users.UserDataInterface {
	return &UserQuery{
		db: db,
	}
}

// Insert implements users.UserDataInterface.
func (repo *UserQuery) Insert(input users.UserCore) error {
	fmt.Println("lead_id before mapping to model:", input.UserLeadID, reflect.TypeOf(input.UserLeadID))
	var userModel = UserCoreToModel(input)
	fmt.Println("lead_id after mapping to model:", input.UserLeadID, reflect.TypeOf(input.UserLeadID))
	fmt.Println(userModel)

	hass, errHass := helper.HassPassword(userModel.Password)
	if errHass != nil {
		return errHass
	}
	userModel.Password = hass
	fmt.Println(userModel)

	tx := repo.db.Create(&userModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}
	return nil
}

func (repo *UserQuery) SelectById(id uint) (users.UserCore, error) {
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

func (repo *UserQuery) Delete(id uint) error {
	tx := repo.db.Where("id = ?", id).Delete(&User{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}
	return nil
}

// SelectAll implements users.UserDataInterface.
func (repo *UserQuery) SelectAll(role_id uint, division_id uint) ([]users.UserCore, error) {
	var userModel []User
	// fmt.Println("divisi_id:", division_id)
	if role_id == 3 || role_id == 4 {
		tx := repo.db.Where("role_id in (3,4) and division_id = ?", division_id).Preload("Role").Preload("Division").Preload("UserImport").Find(&userModel)
		if tx.Error != nil {
			return nil, tx.Error
		}
		if tx.RowsAffected == 0 {
			return nil, errors.New("no row affected")
		}
	} else if role_id == 2 {
		tx := repo.db.Where("role_id in (3,4)").Preload("Role").Preload("Division").Preload("UserImport").Find(&userModel)
		if tx.Error != nil {
			return nil, tx.Error
		}
		if tx.RowsAffected == 0 {
			return nil, errors.New("no row affected")
		}
	} else if role_id == 1 {
		tx := repo.db.Where("role_id in (1,2,3,4)").Preload("Role").Preload("Division").Preload("UserImport").Find(&userModel)
		if tx.Error != nil {
			return nil, tx.Error
		}
		if tx.RowsAffected == 0 {
			return nil, errors.New("no row affected")
		}
	}
	// fmt.Println("ini id lead:", userModel[0].UserLeadID)
	// fmt.Println("ini id dari user lead:", userModel[0].UserLead.ID)

	var userCore []users.UserCore

	for _, value := range userModel {
		var user = ModelToCore(value)
		userCore = append(userCore, user)
	}

	return userCore, nil
}

// Update implements users.UserDataInterface.
func (repo *UserQuery) Update(id uint, input users.UserCore) error {
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

func (repo *UserQuery) Login(email string, password string) (dataLogin users.UserCore, err error) {

	var data User

	tx := repo.db.Where("email = ?", email).Preload("Role").Preload("Division").Find(&data)
	if tx.Error != nil {
		return users.UserCore{}, tx.Error
	}
	check := helper.CheckPassword(password, data.Password)
	if !check {
		return users.UserCore{}, errors.New("password incorect")
	}
	if tx.RowsAffected == 0 {
		return users.UserCore{}, errors.New("data not found")
	}
	dataLogin = ModelToCore(data)
	repo.dataLogin = dataLogin
	return dataLogin, nil
}
