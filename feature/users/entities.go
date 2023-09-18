package users

import (
	divisionCore "hris-app-golang/feature/divisions"
	roleCore "hris-app-golang/feature/roles"
	"time"
)

type UserCore struct {
	ID           string
	FirstName    string
	LastName     string
	Email        string
	PhoneNumber  string
	Password     string
	Address      string
	ProfilePhoto string
	UserLeadID   *string
	RoleID       string
	DivisionID   string
	UserLead     *UserCore
	Role         roleCore.RoleCore
	Division     divisionCore.DivisionCore
	UserImport   UserImportantData
	UserEdu      []UserEducationData
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserImportantData struct {
	UserID          string
	BirthPlace      string
	BirthDate       time.Time
	EmergencyName   string
	EmergencyStatus string
	EmergencyPhone  string
	Npwp            string
	Bpjs            string
	Religion        string
	Gender          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type UserEducationData struct {
	UserID       string
	Name         string
	StartYear    string
	GraduateYear string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserDataInterface interface {
	SelectById(id string) (UserCore, error)
	DeleteById(id string) error
	Login(email string, password string) (dataLogin UserCore, err error)
}

type UserServiceInterface interface {
	GetById(id string) (UserCore, error)
	DeleteUserById(id string) error
	LoginUser(email string, password string) (dataLogin UserCore, token string, err error)
}
