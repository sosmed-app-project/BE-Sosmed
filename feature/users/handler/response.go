package handler

import "time"

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
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserImportantDataResponse struct {
	ID          string    `json:"id"`
	Birth_Place string    `json:"birth_place"`
	Birth_Date  time.Time `json:"birth_date"`
	Religion    string    `json:"Religion"`
}
