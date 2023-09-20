package handler

import (
	"hris-app-golang/feature/divisions"
	roles "hris-app-golang/feature/roles"
	"hris-app-golang/feature/users"
	"time"
)

type LoginResponse struct {
	ID       uint   `json:"id"`
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

type UserResponse struct {
	ID            uint                      `json:"id"`
	FirstName     string                    `json:"first_name"`
	LastName      string                    `json:"last_name"`
	Email         string                    `json:"email"`
	PhoneNumber   string                    `json:"phone_number"`
	Address       string                    `json:"address"`
	Division      DivisionResponse          `json:"division"`
	Role          RoleResponse              `json:"role"`
	UserImportant UserImportantDataResponse `json:"user_important_data"`
}

type DivisionResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type RoleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserImportantDataResponse struct {
	ID          uint   `json:"id"`
	Birth_Place string `json:"birth_place"`
	Birth_Date  string `json:"birth_date"`
	Religion    string `json:"Religion"`
}

type ManagerResponse struct {
	ID        uint
	FirstName string
	LastName  string
	Division  string
}

func UserCoreToResponse(input users.UserCore) UserResponse {
	var resultResponse = UserResponse{
		ID:            input.ID,
		FirstName:     input.FirstName,
		LastName:      input.LastName,
		Email:         input.Email,
		PhoneNumber:   input.PhoneNumber,
		Address:       input.Address,
		Division:      DivisionCoreToResponse(input.Division),
		Role:          RoleCoreToResp(input.Role),
		UserImportant: UserImportCoreToResponse(input.UserImport),
	}
	return resultResponse
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

func UserImportCoreToResponse(input users.UserImportantData) UserImportantDataResponse {
	var resp = UserImportantDataResponse{
		ID:          input.ID,
		Birth_Place: input.BirthPlace,
		Birth_Date:  input.BirthDate,
		Religion:    input.Religion,
	}
	return resp
}

func UserCoreToManagerResponse(input users.UserCore) ManagerResponse {
	var userMan = ManagerResponse{
		ID:        input.ID,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Division:  input.Division.Name,
	}
	return userMan
}
