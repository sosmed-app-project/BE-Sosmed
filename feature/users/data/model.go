package data

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           string
	FirstName    string
	LastName     string
	Email        string
	PhoneNumber  string
	Password     string
	Address      string
	ProfilePhoto string
	UserLeadID   string
	RoleID       string
	DivisionID   string
	UserImport   UserImportant
	UserEdu      UserEducation
	CreatedAt    time.Time
	UpdatedAt    time.Time
	gorm.DeletedAt
}

type UserImportant struct {
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
	gorm.DeletedAt
}

type UserEducation struct {
	UserID       string
	Name         string
	StartYear    string
	GraduateYear string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	gorm.DeletedAt
}
