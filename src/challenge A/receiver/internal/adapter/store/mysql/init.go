package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

type MySQLStore struct {
	db *gorm.DB
}

func InitMySql(dsn string) MySQLStore {
	database, err := gorm.Open(mysql.Open(strings.TrimSpace(dsn)), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	if aErr := database.AutoMigrate(&Order{}); aErr != nil {
		panic("Failed to auto migrate database!")
	}

	return MySQLStore{db: database}
}
