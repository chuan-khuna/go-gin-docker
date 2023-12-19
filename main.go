package main

import (
	"fmt"
	"gin-app/config"
	"gin-app/models"
	"gin-app/routers"

	"github.com/gin-gonic/gin"
)

// import "gin-app/routers": means
// this project(go.mod) > routers folder
// (all files in routers folder start with `package routers`)
// by convention, package name == folder name (we import path/to/folder)

func main() {

	// ------------------------------
	// Register Routers
	// ------------------------------
	r := gin.New()
	r.Use(gin.Logger())

	models.GetDB()
	models.MakeMigrations()

	apiGroup := r.Group("/api")
	{
		// use /api/... to access routers
		routers.HelloRouter(apiGroup)
		routers.PingRouter(apiGroup)
		routers.BookRouter(apiGroup)
	}

	// ------------------------------
	// Run Server
	// ------------------------------
	PORT := config.Port

	fmt.Println("\n*********************************\n")
	fmt.Println("🥃\tServer Running on Port: " + PORT)
	fmt.Println("\n*********************************\n")

	r.Run(":" + PORT)
}
