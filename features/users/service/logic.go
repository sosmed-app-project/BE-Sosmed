package service

import (
	"app-sosmed/app/middlewares"
	"app-sosmed/features/users"
)

type UserService struct {
	userData users.UserDataInterface
}

func NewUsersLogic(repo users.UserDataInterface) users.UserServiceInterface {
	return &UserService{
		userData: repo,
	}
}

func (service *UserService) LoginService(email string, password string) (dataLogin users.UserCore, token string, err error) {
	dataLogin, err = service.userData.LoginQuery(email, password)
	if err != nil {
		return users.UserCore{}, "", err
	}
	token, err = middlewares.CreateToken(dataLogin.ID)
	if err != nil {
		return users.UserCore{}, "", err
	}
	return dataLogin, token, nil
}

func (service *UserService) Create(input users.UserCore) error {
	return service.userData.Insert(input)
}

func (service *UserService) Delete(id uint) error {
	return service.userData.Delete(id)
}

func (service *UserService) Get(ID uint) ([]users.UserCore, error) {
	return service.userData.GetProfile(ID)
}

func (service *UserService) Update(id uint, input users.UserCore) error {
	return service.userData.UpdateProfile(id, input)
}
