package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasvmiguel/go-analytics/controllers"
)

func Start(version string, port string) {
	r := gin.Default()
	r.POST("/"+version+"/notification", controllers.NotificationController)
	r.GET("/"+version+"/ws", controllers.WebsocketController)
	r.Run(port)
}
