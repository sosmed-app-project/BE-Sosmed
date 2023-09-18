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

type UserResponse struct {
	ID                  string                    `json:"id"`
	First_Name          string                    `json:"first_name"`
	Last_Name           string                    `json:"last_name"`
	Email               string                    `json:"email"`
	Phone_Number        string                    `json:"phone_number"`
	Address             string                    `json:"address"`
	Division            DivisionResponse          `json:"division"`
	Role                RoleResponse              `json:"role"`
	User_Important_Data UserImportantDataResponse `json:"user_important_data"`
}

type DivisionResponse struct {
	ID   string `json:"id"`
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

type UserImportantDataResponse struct {
	ID          string    `json:"id"`
	Birth_Place string    `json:"birth_place"`
	Birth_Date  time.Time `json:"birth_date"`
	Religion    string    `json:"Religion"`
}
