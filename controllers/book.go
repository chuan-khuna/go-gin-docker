package controllers

import (
	"net/http"

	"gin-app/models"

	"github.com/gin-gonic/gin"
)

func ListBooks(c *gin.Context) {
	var books []models.Book

	db := models.GetDB()
	db.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}
