package main

import (
	"app-sosmed/app/config"
	"app-sosmed/app/database"
	"fmt"

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
}
