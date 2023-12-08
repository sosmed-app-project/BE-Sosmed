package data

import (
	"app-sosmed/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"column:username;not null;unique"`
	TanggalLahir string `gorm:"column:tanggallahir;not null"`
	Email        string `gorm:"column:email;unique"`
	NoHandphone  string `gorm:"column:nohandphone;not null"`
	Password     string `gorm:"column:password"`
}

func UserCoreToModel(input users.UserCore) User {
	var userModel = User{
		Model:        gorm.Model{},
		Username:     input.Username,
		TanggalLahir: input.TanggalLahir,
		Email:        input.Email,
		NoHandphone:  input.NoHandphone,
		Password:     input.Password,
	}
	return userModel
}

func UserModelToCore(input User) users.UserCore {
	var userCore = users.UserCore{
		ID:           input.ID,
		Username:     input.Username,
		TanggalLahir: input.TanggalLahir,
		Email:        input.Email,
		NoHandphone:  input.NoHandphone,
		Password:     input.Password,
	}
	return userCore
}
