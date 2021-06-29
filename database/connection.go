package database

import (
	"github.com/rajeshj3/jwt-auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Global database cursor
var DB *gorm.DB
var err error

func Connect() {
	// Connect to PostgreSQL database
	dsn := "host=db user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("[DATABASE] Connection Fail")
	}

	// Create tables
	DB.AutoMigrate(&models.User{})
}
