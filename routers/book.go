package routers

import (
	"gin-app/controllers"

	"github.com/gin-gonic/gin"
)

func BookRouter(g *gin.RouterGroup) {
	g.GET("/books", controllers.ListBooks)
}
