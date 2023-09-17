package data

import (
	divisionModel "hris-app-golang/feature/divisions/data"
	roleModel "hris-app-golang/feature/roles/data"
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
	UserLeadID   *string                `gorm:"column:user_lead_id"`
	RoleID       string                 `gorm:"column:role_id;type:varchar(30)"`
	DivisionID   string                 `gorm:"column:division_id;type:varchar(30)"`
	UserLead     *User                  `gorm:"foreignKey:UserLeadID;type:varchar(30)"`
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
