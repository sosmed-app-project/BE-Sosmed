package handler

import (
	"hris-app-golang/feature/divisions"
	roles "hris-app-golang/feature/roles"
	"hris-app-golang/feature/users"
	"time"
)

type LoginResponse struct {
	Role     string `json:"role"`
	Division string `json:"division"`
	Token    string `json:"token"`
}

type UserResponseAll struct {
	ID        uint             `json:"id"`
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	Email     string           `json:"email"`
	Gender    string           `json:"gender"`
	Division  DivisionResponse `json:"division"`
	Role      RoleResponse     `json:"role"`
	CreatedAt time.Time        `json:"created_at"`
}

type DivisionResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type RoleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func UserCoreToResponseAll(input users.UserCore) UserResponseAll {
	var userResp = UserResponseAll{
		ID:        input.ID,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Gender:    input.UserImport.Gender,
		Division:  DivisionCoreToResponse(input.Division),
		Role:      RoleCoreToResp(input.Role),
		CreatedAt: input.CreatedAt,
	}
	return userResp
}

func DivisionCoreToResponse(input divisions.DivisionCore) DivisionResponse {
	var divResp = DivisionResponse{
		ID:   input.ID,
		Name: input.Name,
	}
	return divResp
}

func RoleCoreToResp(input roles.RoleCore) RoleResponse {
	var role = RoleResponse{
		ID:   input.ID,
		Name: input.Name,
	}
	return role
}
