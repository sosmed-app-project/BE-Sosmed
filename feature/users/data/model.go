package data

import (
	"hris-app-golang/feature/divisions"
	divisionModel "hris-app-golang/feature/divisions/data"
	roleCore "hris-app-golang/feature/roles"
	roleModel "hris-app-golang/feature/roles/data"
	"hris-app-golang/feature/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint                   `gorm:"column:id;primaryKey"`
	FirstName    string                 `gorm:"column:first_name;not null"`
	LastName     string                 `gorm:"column:last_name;not null"`
	Email        string                 `gorm:"column:email;not null;unique"`
	PhoneNumber  string                 `gorm:"column:phone_number;unique"`
	Password     string                 `gorm:"column:password;not null;default:'qwerty'"`
	Address      string                 `gorm:"column:address"`
	ProfilePhoto string                 `gorm:"column:profile_photo"`
	UserLeadID   uint                   `gorm:"column:user_lead_id"`
	RoleID       *uint                  `gorm:"column:role_id"`
	DivisionID   *uint                  `gorm:"column:division_id"`
	Role         roleModel.Role         `gorm:"foreignKey:RoleID"`
	Division     divisionModel.Division `gorm:"foreignKey:DivisionID"`
	UserImport   UserImportant          `gorm:"foreignKey:UserID"`
	UserEdu      []UserEducation        `gorm:"foreignKey:UserID"`
	CreatedAt    time.Time              `gorm:"column:created_at"`
	UpdatedAt    time.Time              `gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt         `gorm:"column:deleted_at;index"`
	// UserLead     *User                  `gorm:"foreignKey:UserLeadID"`
}

type UserImportant struct {
	ID              uint           `gorm:"column:id;primaryKey"`
	UserID          uint           `gorm:"column:user_id"`
	BirthPlace      string         `gorm:"column:birth_place"`
	BirthDate       time.Time      `gorm:"column:birth_date"`
	EmergencyName   string         `gorm:"column:emergency_name"`
	EmergencyStatus string         `gorm:"column:emergency_status"`
	EmergencyPhone  string         `gorm:"column:emergency_phone"`
	Npwp            string         `gorm:"column:npwp"`
	Bpjs            string         `gorm:"column:bpjs"`
	Religion        string         `gorm:"type:enum('Islam','Kristen','Katolik','Hindu','Budha','Konghucu')"`
	Gender          string         `gorm:"type:enum('Male','Female')"`
	CreatedAt       time.Time      `gorm:"column:created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

type UserEducation struct {
	ID           uint           `gorm:"primaryKey"`
	UserID       uint           `gorm:"column:user_id"`
	Name         string         `gorm:"column:name"`
	StartYear    string         `gorm:"column:start_year"`
	GraduateYear string         `gorm:"column:graduate_year"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func UserCoreToModel(input users.UserCore) User {
	var userModel = User{
		ID:           input.ID,
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		Password:     input.Password,							
		Address:      input.Address,
		ProfilePhoto: input.ProfilePhoto,
		UserLeadID:   input.UserLeadID,
		RoleID:       &input.RoleID,
		DivisionID:   &input.DivisionID,
		UserImport:   UserImportantCoreToModel(input.UserImport),
		UserEdu:      UserEducationCoreToModel(input.UserEdu),
		// UserLead:     UserLeadCoreToModel(input.UserLead),
	}
	return userModel
}

// func UserLeadCoreToModel(input *users.UserCore) *User {
// 	var userLead User
// 	if input != nil {
// 		userLead = User{
// 			ID:        input.ID,
// 			FirstName: input.FirstName,
// 			LastName:  input.LastName,
// 			Email:     input.Email,
// 			CreatedAt: input.CreatedAt,
// 			UpdatedAt: input.UpdatedAt,
// 		}
// 	}
// 	return &userLead
// }

func UserImportantCoreToModel(input users.UserImportantData) UserImportant {
	var userImport = UserImportant{
		BirthPlace:      input.BirthPlace,
		BirthDate:       input.BirthDate,
		EmergencyName:   input.EmergencyName,
		EmergencyStatus: input.EmergencyPhone,
		EmergencyPhone:  input.EmergencyStatus,
		Npwp:            input.Npwp,
		Bpjs:            input.Bpjs,
		Religion:        input.Religion,
		Gender:          input.Gender,
	}
	return userImport
}

func UserEducationCoreToModel(input []users.UserEducationData) []UserEducation {
	var userEduList []UserEducation
	for _, value := range input {
		var userEdu = UserEducation{
			Name:         value.Name,
			StartYear:    value.StartYear,
			GraduateYear: value.GraduateYear,
		}
		userEduList = append(userEduList, userEdu)
	}
	return userEduList
}

func ModelToCore(input User) users.UserCore {
	var userCore = users.UserCore{
		ID:           input.ID,
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		Password:     input.Password,
		Address:      input.Address,
		ProfilePhoto: input.ProfilePhoto,
		UserLeadID:   input.UserLeadID,
		RoleID:       *input.RoleID,
		DivisionID:   *input.DivisionID,
		Role:         RoleModelToCore(input.Role),
		Division:     DivisionModelToCore(input.Division),
		UserImport:   UserImportModelToCore(input.UserImport),
		UserEdu:      UserEducationModelToCore(input.UserEdu),
		CreatedAt:    input.CreatedAt,
		UpdatedAt:    input.UpdatedAt,
		// UserLead:     UserLeadModelToCore(input.UserLead),
	}
	return userCore
}

// func UserLeadModelToCore(input *User) *users.UserCore {
// 	var userLead users.UserCore
// 	if input != nil {
// 		userLead = users.UserCore{
// 			ID:          input.ID,
// 			FirstName:   input.FirstName,
// 			LastName:    input.LastName,
// 			Email:       input.Email,
// 			PhoneNumber: input.PhoneNumber,
// 			Password:    input.Password,
// 			Address:     input.Address,
// 			RoleID:      *input.RoleID,
// 			DivisionID:  *input.DivisionID,
// 			CreatedAt:   input.CreatedAt,
// 			UpdatedAt:   input.UpdatedAt,
// 		}
// 	}
// 	return &userLead
// }

func RoleModelToCore(input roleModel.Role) roleCore.RoleCore {
	var role = roleCore.RoleCore{
		ID:        input.ID,
		Name:      input.Name,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}
	return role
}

func DivisionModelToCore(input divisionModel.Division) divisions.DivisionCore {
	var division = divisions.DivisionCore{
		ID:        input.ID,
		Name:      input.Name,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}
	return division
}

func UserImportModelToCore(input UserImportant) users.UserImportantData {
	var userImport = users.UserImportantData{
		ID:              input.ID,
		UserID:          input.UserID,
		BirthPlace:      input.BirthPlace,
		BirthDate:       input.BirthDate,
		EmergencyName:   input.EmergencyName,
		EmergencyStatus: input.EmergencyStatus,
		EmergencyPhone:  input.EmergencyPhone,
		Npwp:            input.Npwp,
		Bpjs:            input.Bpjs,
		Religion:        input.Religion,
		Gender:          input.Gender,
		CreatedAt:       input.CreatedAt,
		UpdatedAt:       input.UpdatedAt,
	}
	return userImport
}

func UserEducationModelToCore(input []UserEducation) []users.UserEducationData {
	var userEduList []users.UserEducationData
	for _, value := range input {
		var userEdu = users.UserEducationData{
			ID:           value.ID,
			UserID:       value.UserID,
			Name:         value.Name,
			StartYear:    value.StartYear,
			GraduateYear: value.GraduateYear,
			CreatedAt:    value.CreatedAt,
			UpdatedAt:    value.UpdatedAt,
		}
		userEduList = append(userEduList, userEdu)
	}
	return userEduList
}
