package users

import "time"

type UserCore struct {
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
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UserImport   UserImportantData
	UserEdu      UserEducationData
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
