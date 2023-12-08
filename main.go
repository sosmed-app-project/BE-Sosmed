// main.go
package main

import (
	"app-sosmed/app/config"
	"app-sosmed/app/database"
	"app-sosmed/app/router"
	"app-sosmed/feature/users/data"
	"app-sosmed/feature/users/handler"
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func main() {
	config := config.InitConfig()

	if config == nil {
		fmt.Println("Failed to read config")
		return
	}

	db, err := database.InitMySQL(*config)
	if err != nil {
		fmt.Println("Failed to connect to the database. Error:", err)
		return
	}
	defer func(db *gorm.DB) {
		sqlDB, err := db.DB()
		if err != nil {
			fmt.Println("Failed to get SQL DB. Error:", err)
			return
		}
		sqlDB.Close()
	}(db)

	fmt.Println("Connected to the database successfully!")

	e := echo.New()

	userRepository := data.NewUserRepository(db)
	userHandler := handler.NewUserHandler(userRepository)

	router.InitRoute(e, userHandler)

	e.Start(":8000")
}
