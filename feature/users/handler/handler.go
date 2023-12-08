// feature/users/handler/handler.go
package handler

import (
	"app-sosmed/app/auth"
	"app-sosmed/feature/users/data"
	"app-sosmed/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	UserRepo *data.UserRepository
	Token    string // Added field to store token
}

// NewUserHandler creates a new user handler
func NewUserHandler(userRepo *data.UserRepository) *UserHandler {
	return &UserHandler{
		UserRepo: userRepo,
	}
}

// Register handles user registration
func (h *UserHandler) Register(c echo.Context) error {
	var user data.User
	if err := c.Bind(&user); err != nil {
		return helper.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
	}

	// Validate input data (ensure TanggalLahir and NoHandphone are not empty)
	if user.TanggalLahir == "" || user.NoHandphone == "" {
		return helper.RespondWithError(c, http.StatusBadRequest, "Tanggal lahir dan nomor telepon harus diisi")
	}

	// Hash the password before saving it to the database
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return helper.RespondWithError(c, http.StatusInternalServerError, "Failed to hash password")
	}
	user.Password = hashedPassword

	err = h.UserRepo.CreateUser(&user)
	if err != nil {
		return helper.RespondWithError(c, http.StatusInternalServerError, "Failed to create user")
	}

	// Generate JWT token
	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return helper.RespondWithError(c, http.StatusInternalServerError, "Failed to generate JWT token")
	}

	// Respond with success message, user data, and token
	response := map[string]interface{}{
		"message": "Registration successful",
		"user": map[string]interface{}{
			"ID":           user.ID,
			"Username":     user.Username,
			"TanggalLahir": user.TanggalLahir,
			"NoHandphone":  user.NoHandphone,
			"Email":        user.Email,
			"CreatedAt":    user.CreatedAt,
			"UpdatedAt":    user.UpdatedAt,
			"DeletedAt":    user.DeletedAt,
		},
		"token": token,
	}

	return helper.RespondWithJSON(c, http.StatusCreated, response)
}

// Login handles user login
func (h *UserHandler) Login(c echo.Context) error {
	var user data.User
	if err := c.Bind(&user); err != nil {
		return helper.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
	}

	// Retrieve the user from the database based on the username
	retrievedUser, err := h.UserRepo.GetUserByUsername(user.Username)
	if err != nil {
		return helper.RespondWithError(c, http.StatusUnauthorized, "Invalid username or password")
	}

	// Check if the provided password matches the stored hashed password
	if !auth.CheckPasswordHash(user.Password, retrievedUser.Password) {
		return helper.RespondWithError(c, http.StatusUnauthorized, "Invalid username or password")
	}

	// Generate JWT token
	token, err := auth.GenerateToken(retrievedUser.ID)
	if err != nil {
		return helper.RespondWithError(c, http.StatusInternalServerError, "Failed to generate JWT token")
	}

	// Set the generated token to the Token field in the UserHandler
	h.Token = token

	// Respond with success message and token
	response := map[string]interface{}{
		"message": "Login successful",
		"token":   token,
	}

	return helper.RespondWithJSON(c, http.StatusOK, response)
}
