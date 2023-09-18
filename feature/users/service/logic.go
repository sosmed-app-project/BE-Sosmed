package service

import "hris-app-golang/feature/users"

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

// Delete implements users.UserServiceInterface.
func (*UserService) Delete(id string) error {
	panic("unimplemented")
}

// GetAll implements users.UserServiceInterface.
func (service *UserService) GetAll(role_id string, division_id string) ([]users.UserCore, error) {
	result, err := service.userData.SelectAll(role_id, division_id)
	return result, err
}

// GetById implements users.UserServiceInterface.
func (*UserService) GetById(id string) (users.UserCore, error) {
	panic("unimplemented")
}

// Login implements users.UserServiceInterface.
func (service *UserService) Login(email string, password string) (users.UserCore, string, error) {
	result, err := service.userData.Login(email, password)

	var token string

	return result, token, err
}

// Update implements users.UserServiceInterface.
func (service *UserService) Update(id string, input users.UserCore) error {
	err := service.userData.Update(id, input)
	return err
}
