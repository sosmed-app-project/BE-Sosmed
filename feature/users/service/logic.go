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
func (service *UserService) GetAll(role_id, division_id, page, item uint, search_name string) ([]users.UserCore, bool, error) {
	result, count, err := service.userData.SelectAll(role_id, division_id, page, item, search_name)

	next := true
	var pages int64
	if item != 0 {
		pages = count / int64(item)
		if count%int64(item) != 0 {
			pages += 1
		}
		if page == uint(pages) {
			next = false
		}
	}

	return result, next, err
}

// Update implements users.UserServiceInterface.
func (service *UserService) Update(id uint, input users.UserCore) error {
	err := service.userData.Update(id, input)
	return err
}

func (service *UserService) GetById(id uint) (users.UserCore, error) {
	return service.userData.SelectById(id)
}

func (service *UserService) Delete(id uint) error {
	err := service.userData.Delete(id)
	return err
}

func (service *UserService) Login(email string, password string) (dataLogin users.UserCore, token string, err error) {

	dataLogin, err = service.userData.Login(email, password)
	if err != nil {
		return users.UserCore{}, "", err
	}
	token, err = middlewares.CreateToken(dataLogin.ID, dataLogin.RoleID, dataLogin.DivisionID)
	if err != nil {
		return users.UserCore{}, "", err
	}
	return dataLogin, token, nil
}
