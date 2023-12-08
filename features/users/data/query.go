package data

import (
	"app-sosmed/features/users"
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserQuery struct {
	db        *gorm.DB
	dataLogin users.UserCore
}

func NewUsersQuery(db *gorm.DB) users.UserDataInterface {
	return &UserQuery{
		db: db,
	}
}

func (repo *UserQuery) LoginQuery(email string, password string) (dataLogin users.UserCore, err error) {

	var data User

	tx := repo.db.Where("email = ? && password = ?", email, password).Find(&data)
	if tx.Error != nil {
		return users.UserCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return users.UserCore{}, errors.New("no row affected")
	}
	dataLogin = UserModelToCore(data)
	repo.dataLogin = dataLogin
	return dataLogin, nil
}

func (repo *UserQuery) Insert(input users.UserCore) error {

	var userModel = UserCoreToModel(input)

	tx := repo.db.Create(&userModel)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
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

func (repo *UserQuery) GetProfile(ID uint) ([]users.UserCore, error) {
	var userData []User

	query := repo.db.Where("id = ?", ID)
	tx := query.Find(&userData)

	if tx.Error != nil {
		return nil, tx.Error
	}

	var userCore []users.UserCore
	for _, value := range userData {
		userCore = append(userCore, users.UserCore{
			ID:           value.ID,
			Username:     value.Username,
			TanggalLahir: value.TanggalLahir,
			Email:        value.Email,
			NoHandphone:  value.NoHandphone,
			Password:     value.Password,
			CreatedAt:    time.Time{},
			UpdatedAt:    time.Time{},
			DeletedAt:    time.Time{},
		})
	}

	return userCore, nil
}

func (repo *UserQuery) UpdateProfile(id uint, input users.UserCore) error {
	var user User
	tx := repo.db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("target not found")
	}

	updatedUser := UserCoreToModel(input)

	tx = repo.db.Model(&user).Updates(updatedUser)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + " failed to update data")
	}
	return nil
}
