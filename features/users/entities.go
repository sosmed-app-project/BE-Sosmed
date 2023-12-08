package users

import (
	"time"
)

type UserCore struct {
	ID           uint
	Username     string
	TanggalLahir string
	Email        string
	NoHandphone  string
	Password     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type UserDataInterface interface {
	LoginQuery(email, password string) (UserCore, error)
	Insert(input UserCore) error
	Delete(id uint) error
	GetProfile(ID uint) ([]UserCore, error)
	UpdateProfile(id uint, input UserCore) error
}

type UserServiceInterface interface {
	LoginService(email, password string) (UserCore, string, error)
	Create(input UserCore) error
	Delete(id uint) error
	Get(ID uint) ([]UserCore, error)
	Update(id uint, input UserCore) error
}
