package handler

import (
	"hris-app-golang/feature/users"
	"time"
)

type UserRequest struct {
	FirstName    string                 `json:"first_name" form:"first_name"`
	LastName     string                 `json:"last_name" form:"last_name"`
	Email        string                 `json:"email" form:"email"`
	PhoneNumber  string                 `json:"phone_number" form:"phone_number"`
	Password     string                 `json:"password" form:"password"`
	Address      string                 `json:"address" form:"address"`
	ProfilePhoto string                 `json:"profile_photo" form:"profile_photo"`
	UserLeadID   uint                   `json:"user_lead_id" form:"user_lead_id"`
	RoleID       uint                   `json:"role_id" form:"role_id"`
	DivisionID   uint                   `json:"division_id" form:"division_id"`
	UserImport   UserImportantRequest   `json:"user_important_data" form:"user_important_data"`
	UserEdu      []UserEducationRequest `json:"user_education_data" form:"user_education_data"`
}

type UserImportantRequest struct {
	BirthPlace      string    `json:"birth_place" form:"birth_place"`
	BirthDate       time.Time `json:"birth_date" form:"birth_date"`
	EmergencyName   string    `json:"emergency_name" form:"emergency_name"`
	EmergencyStatus string    `json:"emergency_status" form:"emergency_status"`
	EmergencyPhone  string    `json:"emergency_phone" form:"emergency_phone"`
	Npwp            string    `json:"npwp" form:"npwp"`
	Bpjs            string    `json:"bpjs" form:"bpjs"`
	Religion        string    `json:"religion" form:"religion"`
	Gender          string    `json:"gender" form:"gender"`
}

type UserEducationRequest struct {
	Name         string `json:"name" form:"name"`
	StartYear    string `json:"start_year" form:"start_year"`
	GraduateYear string `json:"graduate_year" form:"graduate_year"`
}

func UserRequestToCore(input UserRequest) users.UserCore {
	var userCore = users.UserCore{
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		Password:     input.Password,
		Address:      input.Address,
		ProfilePhoto: input.ProfilePhoto,
		UserLeadID:   &input.UserLeadID,
		RoleID:       input.RoleID,
		DivisionID:   input.DivisionID,
		UserImport:   UserImportantRequestToCore(input.UserImport),
		UserEdu:      UserEducationRequestToCore(input.UserEdu),
	}
	return userCore
}

func UserImportantRequestToCore(input UserImportantRequest) users.UserImportantData {
	var userImport = users.UserImportantData{
		BirthPlace:      input.BirthPlace,
		BirthDate:       input.BirthDate,
		EmergencyName:   input.EmergencyName,
		EmergencyStatus: input.EmergencyStatus,
		EmergencyPhone:  input.EmergencyPhone,
		Npwp:            input.Npwp,
		Bpjs:            input.Bpjs,
		Religion:        input.Religion,
		Gender:          input.Gender,
	}
	return userImport
}

func UserEducationRequestToCore(input []UserEducationRequest) []users.UserEducationData {
	var userEduList []users.UserEducationData
	for _, value := range input {
		var userEdu = users.UserEducationData{
			Name:         value.Name,
			StartYear:    value.StartYear,
			GraduateYear: value.GraduateYear,
		}
		userEduList = append(userEduList, userEdu)
	}
	return userEduList
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
