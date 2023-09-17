package service

import "hris-app-golang/feature/users"

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
