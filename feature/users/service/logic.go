package service

<<<<<<< HEAD
import "hris-app-golang/feature/users"

type UserService struct {
=======
import (
	"hris-app-golang/app/middlewares"
	"hris-app-golang/feature/users"
)

type userService struct {
>>>>>>> d65206d7f3cdd592e18a4c27ff9f11dcaa5ebf9b
	userData users.UserDataInterface
}

func New(repo users.UserDataInterface) users.UserServiceInterface {
<<<<<<< HEAD
	return &UserService{
=======
	return &userService{
>>>>>>> d65206d7f3cdd592e18a4c27ff9f11dcaa5ebf9b
		userData: repo,
	}
}

<<<<<<< HEAD
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
=======
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
>>>>>>> d65206d7f3cdd592e18a4c27ff9f11dcaa5ebf9b
}
