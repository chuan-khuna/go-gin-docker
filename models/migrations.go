package models

import (
	"fmt"
)

func MakeMigrations() {
	db := GetDB()
	db.AutoMigrate(&Book{})
	fmt.Println("\n*********************************\n")
	fmt.Println("🥃\tMigrations Done")
	fmt.Println("\n*********************************\n")
}
