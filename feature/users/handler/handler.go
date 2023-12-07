package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUserHandler mengambil data pengguna berdasarkan ID
func GetUserHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Panggil fungsi untuk mendapatkan data pengguna dari database
	user, err := getUserFromDatabase(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	c.JSON(http.StatusOK, user)
}
