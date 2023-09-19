package handler

import levels "hris-app-golang/feature/roles"

type RoleRequest struct {
	Name string `json:"name"`
}

func RequestToCore(input RoleRequest) levels.RoleCore {
	return levels.RoleCore{
		Name: input.Name,
	}
}
