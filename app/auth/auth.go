// app/auth/auth.go
package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Claims is the struct for JWT claims
type Claims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

// GenerateToken generates a JWT token
func GenerateToken(idUser uint) (string, error) {
	claims := &Claims{
		ID: idUser,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err := token.SignedString([]byte("$!1gnK3yyy!!!"))
	if err != nil {
		return "", err
	}

	return strToken, nil
}

// VerifyToken verifies the JWT token
func VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("$!1gnK3yyy!!!"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// HashPassword hashes the provided password using bcrypt
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPasswordHash checks if the provided password matches the stored hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ExtractUserID extracts user ID from the token claims
func ExtractUserID(c echo.Context) (uint, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*Claims)
	return claims.ID, nil
}
