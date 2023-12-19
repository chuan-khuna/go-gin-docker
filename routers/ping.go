package routers

import (
	"gin-app/controllers"

	"github.com/gin-gonic/gin"
)

func PingRouter(g *gin.RouterGroup) {
	g.GET("/ping", controllers.PingController)
}
