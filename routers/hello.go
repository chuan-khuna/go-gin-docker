package routers

import (
	"gin-app/controllers"

	"github.com/gin-gonic/gin"
)

func HelloRouter(g *gin.RouterGroup) {
	g.GET("/hello", controllers.HelloController)
}
