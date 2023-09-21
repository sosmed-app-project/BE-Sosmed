package service

import (
	"hris-app-golang/feature/roles"
)

type RoleService struct {
	roleData roles.RoleDataInterface
}

func NewRoleService(repo roles.RoleDataInterface) roles.RoleServiceInterface {
	return &RoleService{
		roleData: repo,
	}
}

func (service *RoleService) GetAllRoles() ([]roles.RoleCore, error) {
	result, err := service.roleData.Read()
	if err != nil {
		return []roles.RoleCore{}, err
	}
	return result, err
}
