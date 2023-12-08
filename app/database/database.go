// app/database/database.go
package database

import (
	"app-sosmed/app/config"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitMySQL initializes the MySQL database connection
func InitMySQL(c config.AppConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUSER,
		c.DBPASS,
		c.DBHOST,
		c.DBPORT,
		c.DBNAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("Failed to connect to the database. Error:", err.Error())
		return nil, err
	}

	return db, nil
}
