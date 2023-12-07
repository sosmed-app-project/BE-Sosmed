package posting

import (
	"time"
)

// User model
type User struct {
	ID          uint      `json:"id"`
	Username    string    `json:"username"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
