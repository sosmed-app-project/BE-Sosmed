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
	ID           string                 `gorm:"column:id;type:varchar(30);primaryKey"`
	FirstName    string                 `gorm:"column:first_name;not null"`
	LastName     string                 `gorm:"column:last_name;not null"`
	Email        string                 `gorm:"column:email;not null;unique"`
	PhoneNumber  string                 `gorm:"column:phone_number;unique"`
	Password     string                 `gorm:"column:password;not null;default:'qwerty'"`
	Address      string                 `gorm:"column:address"`
	ProfilePhoto string                 `gorm:"column:profile_photo"`
	UserLeadID   *string                `gorm:"column:user_lead_id;type:varchar(30)"`
	RoleID       string                 `gorm:"column:role_id;type:varchar(30)"`
	DivisionID   string                 `gorm:"column:division_id;type:varchar(30)"`
	UserLead     *User                  `gorm:"foreignKey:UserLeadID"`
	Role         roleModel.Role         `gorm:"foreignKey:RoleID"`
	Division     divisionModel.Division `gorm:"foreignKey:DivisionID"`
	UserImport   UserImportant          `gorm:"foreignKey:UserID"`
	UserEdu      []UserEducation        `gorm:"foreignKey:UserID"`
	CreatedAt    time.Time              `gorm:"column:created_at"`
	UpdatedAt    time.Time              `gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt         `gorm:"column:deleted_at;index"`
}

type UserImportant struct {
	ID              string         `gorm:"column:id;type:varchar(30);primaryKey"`
	UserID          string         `gorm:"column:user_id;type:varchar(30)"`
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
	ID           string         `gorm:"type:varchar(30);primaryKey"`
	UserID       string         `gorm:"column:user_id;type:varchar(30)"`
	Name         string         `gorm:"column:name"`
	StartYear    string         `gorm:"column:start_year"`
	GraduateYear string         `gorm:"column:graduate_year"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func ModelToCore(modelUser User) users.UserCore {
	coreUser := users.UserCore{
		ID:           modelUser.ID,
		FirstName:    modelUser.FirstName,
		LastName:     modelUser.LastName,
		Email:        modelUser.Email,
		PhoneNumber:  modelUser.PhoneNumber,
		Password:     modelUser.Password,
		Address:      modelUser.Address,
		ProfilePhoto: modelUser.ProfilePhoto,
		UserLeadID:   modelUser.UserLeadID,
		RoleID:       modelUser.RoleID,
		DivisionID:   modelUser.DivisionID,
		//UserLead:     &users.UserCore{},
		Role:       RoleModelToCore(modelUser.Role),
		Division:   DivisionModelToCore(modelUser.Division),
		UserImport: UserImportModelToCore(modelUser.UserImport),
		UserEdu:    UserEduModelCore(modelUser.UserEdu),
		CreatedAt:  modelUser.CreatedAt,
		UpdatedAt:  modelUser.UpdatedAt,
	}

	return coreUser
}

func RoleModelToCore(RoleModel roleModel.Role) roleCore.RoleCore {
	role := roleCore.RoleCore{
		ID:        RoleModel.ID,
		Name:      RoleModel.Name,
		CreatedAt: RoleModel.CreatedAt,
		UpdatedAt: RoleModel.UpdatedAt,
	}
	return role
}

func DivisionModelToCore(DivisionModel divisionModel.Division) divisions.DivisionCore {
	divisionsData := divisions.DivisionCore{
		ID:        DivisionModel.ID,
		Name:      DivisionModel.Name,
		CreatedAt: DivisionModel.CreatedAt,
		UpdatedAt: DivisionModel.UpdatedAt,
	}
	return divisionsData
}

func UserImportModelToCore(UserImportModel UserImportant) users.UserImportantData {
	importantCore := users.UserImportantData{
		UserID:          UserImportModel.UserID,
		BirthPlace:      UserImportModel.BirthPlace,
		BirthDate:       UserImportModel.BirthDate,
		EmergencyName:   UserImportModel.EmergencyName,
		EmergencyStatus: UserImportModel.EmergencyStatus,
		EmergencyPhone:  UserImportModel.EmergencyPhone,
		Npwp:            UserImportModel.Npwp,
		Bpjs:            UserImportModel.Bpjs,
		Religion:        UserImportModel.Religion,
		Gender:          UserImportModel.Gender,
		CreatedAt:       UserImportModel.CreatedAt,
		UpdatedAt:       UserImportModel.UpdatedAt,
	}
	return importantCore
}

func UserEduModelCore(UserEduModel []UserEducation) []users.UserEducationData {
	var userEduData []users.UserEducationData

	for _, val := range UserEduModel {
		var userEdu = users.UserEducationData{
			UserID:       val.UserID,
			Name:         val.Name,
			StartYear:    val.StartYear,
			GraduateYear: val.StartYear,
			CreatedAt:    val.CreatedAt,
			UpdatedAt:    val.UpdatedAt,
		}
		userEduData = append(userEduData, userEdu)
	}

	return userEduData
}
