package models

import (
	"fmt"
	"gin-app/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	dbHost := "db"
	dbPort := "5432"
	dbUser := config.DBUser
	dbPassword := config.DBPassword
	dbName := config.DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic("🔥 failed to connect database")
	}

	return db
}

var db *gorm.DB

func GetDB() *gorm.DB {
	db = InitDB()
	return db
}
