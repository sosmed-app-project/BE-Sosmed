package service

import "hris-app-golang/feature/roles"

type RoleService struct {
	RoleData roles.RoleDataInterface
}

// GetAllRoles implements roles.RoleServiceInterface.
func (servis *RoleService) GetAllRoles() ([]roles.RoleCore, error) {
	var res, err = servis.RoleData.GetAllRoles()
	return res, err
}

func New(repo roles.RoleDataInterface) roles.RoleServiceInterface {
	return &RoleService{
		RoleData: repo,
	}
}
