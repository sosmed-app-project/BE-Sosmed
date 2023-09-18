package service

import (
	middlewares "hris-app-golang/app/middlewares"
	"hris-app-golang/feature/users"
)

type UserService struct {
	userData users.UserDataInterface
}

func New(repo users.UserDataInterface) users.UserServiceInterface {
	return &UserService{
		userData: repo,
	}
}

// Add implements users.UserServiceInterface.
func (service *UserService) Add(input users.UserCore) error {
	err := service.userData.Insert(input)
	return err
}

// GetAll implements users.UserServiceInterface.
func (service *UserService) GetAll(role_id string, division_id string) ([]users.UserCore, error) {
	result, err := service.userData.SelectAll(role_id, division_id)
	return result, err
}

// Update implements users.UserServiceInterface.
func (service *UserService) Update(id string, input users.UserCore) error {
	err := service.userData.Update(id, input)
	return err
}

func (service *UserService) GetById(id string) (users.UserCore, error) {
	return service.userData.SelectById(id)
}

func (service *UserService) Delete(id string) error {
	err := service.userData.Delete(id)
	return err
}

func (service *UserService) Login(email string, password string) (dataLogin users.UserCore, token string, err error) {

	dataLogin, err = service.userData.Login(email, password)
	if err != nil {
		return users.UserCore{}, "", err
	}
	token, err = middlewares.CreateToken(dataLogin.ID, dataLogin.Role.Name, dataLogin.Division.Name)
	if err != nil {
		return users.UserCore{}, "", err
	}
	return dataLogin, token, nil
}
