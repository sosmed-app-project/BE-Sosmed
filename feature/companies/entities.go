package companies

import "time"

type CompanyCore struct {
	ID          string
	Name        string
	Address     string
	PhoneNumber string
	Email       string
	Website     string
	Founder     string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
