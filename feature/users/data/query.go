// feature/users/data/query.go
package data

import (
	"gorm.io/gorm"
)

// UserRepository is the repository for user data operations
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// CreateUser creates a new user in the database
func (r *UserRepository) CreateUser(user *User) error {
	return r.DB.Create(user).Error
}

// GetUserByUsername retrieves a user by username from the database
func (r *UserRepository) GetUserByUsername(username string) (*User, error) {
	var user User
	err := r.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}
