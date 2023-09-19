package service

import (
	"fmt"
	levels "hris-app-golang/feature/roles"
)

type RoleService struct {
	roleData levels.RoleDataInterface
}

// UpdateUserById implements levels.RoleServiceInterface.

func New(repo levels.RoleDataInterface) levels.RoleServiceInterface {
	return &RoleService{
		roleData: repo,
	}
}

func (service *RoleService) GetAll(ID uint, Name string) ([]levels.RoleCore, error) {
	result, err := service.roleData.SelectAll(ID, Name)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve roles: %w", err)
	}
	return result, nil
}

func (service *RoleService) UpdateRoleById(ID uint, input levels.RoleCore) (levels.RoleCore, error) {

	result, err := service.roleData.UpdateById(ID, input)
	if err != nil {
		return levels.RoleCore{}, err
	}
	return result, nil
}
