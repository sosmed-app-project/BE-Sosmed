package service

import (
	levels "hris-app-golang/feature/roles"
)

type RoleService struct {
	roleData levels.RoleDataInterface
}

func New(repo levels.RoleDataInterface) levels.RoleServiceInterface {
	return &RoleService{
		roleData: repo,
	}
}

func (service *RoleService) GetAll() ([]levels.RoleCore, error) {
	result, err := service.roleData.SelectAll()
	if err != nil {
		return []levels.RoleCore{}, err
	}
	return result, err
}

func (service *RoleService) UpdateRoleById(ID uint, input levels.RoleCore) (levels.RoleCore, error) {

	result, err := service.roleData.UpdateById(ID, input)
	if err != nil {
		return levels.RoleCore{}, err
	}
	return result, nil
}
