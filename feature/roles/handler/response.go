package handler

import (
	levels "hris-app-golang/feature/roles"
)

type RoleResponseAll struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type RoleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func RoleCoreToResponseAll(input levels.RoleCore) RoleResponseAll {
	var roleResp = RoleResponseAll{
		ID:   input.ID,
		Name: input.Name,
	}
	return roleResp
}
