package config

import "os"

var Port = getPort()

var DB = getDB()
var DBUser = getUser()
var DBPassword = getPassword()

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		// Default to 8200 if PORT is not set
		return "8200"
	}
	return port
}

func getDB() string {
	db := os.Getenv("POSTGRES_NAME")
	return db
}

func getUser() string {
	user := os.Getenv("POSTGRES_USER")
	return user
}

func getPassword() string {
	password := os.Getenv("POSTGRES_PASSWORD")
	return password
}
