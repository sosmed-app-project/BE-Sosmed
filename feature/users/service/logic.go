package service

import (
	"hris-app-golang/app/middlewares"
	"hris-app-golang/feature/users"
)

type userService struct {
	userData users.UserDataInterface
}

func New(repo users.UserDataInterface) users.UserServiceInterface {
	return &userService{
		userData: repo,
	}
}

func (service *userService) GetById(id string) (users.UserCore, error) {
	return service.userData.SelectById(id)
}

func (service *userService) DeleteUserById(id string) error {
	err := service.userData.DeleteById(id)
	return err
}

func (service *userService) LoginUser(email string, password string) (dataLogin users.UserCore, token string, err error) {

	dataLogin, err = service.userData.Login(email, password)
	if err != nil {
		return users.UserCore{}, "", err
	}
	token, err = middlewares.CreateToken(dataLogin.Role.Name, dataLogin.ID)
	if err != nil {
		return users.UserCore{}, "", err
	}
	return dataLogin, token, nil
}
