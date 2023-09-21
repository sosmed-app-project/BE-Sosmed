package service

import (
	middlewares "hris-app-golang/app/middlewares"
	"hris-app-golang/feature/users"
	"mime/multipart"
)

type UserService struct {
	userData users.UserDataInterface
}

func New(repo users.UserDataInterface) users.UserServiceInterface {
	return &UserService{
		userData: repo,
	}
}

// GetAllManager implements users.UserServiceInterface.
func (service *UserService) GetAllManager() ([]users.UserCore, error) {
	result, err := service.userData.GetAllManager()
	return result, err
}

// Add implements users.UserServiceInterface.
func (service *UserService) Add(input users.UserCore, file multipart.File, header *multipart.FileHeader) error {
	err := service.userData.Insert(input, file, header)
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

// CountEmployees 
func (service *UserService) CountEmployees() (uint, error) {
	count, err := service.userData.CountEmployees()
	return count, err
}

// CountManagers 
func (service *UserService) CountManagers() (uint, error) {
	count, err := service.userData.CountManagers()
	return count, err
}

// CountMaleUsers 
func (service *UserService) CountMaleUsers() (uint, error) {
	count, err := service.userData.CountMaleUsers()
	return count, err
}

// CountFemaleUsers 
func (service *UserService) CountFemaleUsers() (uint, error) {
	count, err := service.userData.CountFemaleUsers()
	return count, err
}