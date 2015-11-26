package api

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/lucasvmiguel/go-analytics/controllers"
)

func Start(debugger bool, recovery bool, version string, port string) {
	r := gin.New()

	if debugger {
		r.Use(gin.Logger())
	}

	if recovery {
		r.Use(gin.Recovery())
	}

	r.POST("/"+version+"/auth", controllers.AccessController)
	r.Use(controllers.AuthentificationController)
	r.POST("/"+version+"/notification", controllers.NotificationController)
	r.GET("/"+version+"/ws", controllers.WebsocketController)

	err := r.Run(port)
	if err != nil {
		logrus.Panic("Error to open api")
	}
}
