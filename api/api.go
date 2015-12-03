package api

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/lucasvmiguel/go-analytics/controllers"
	"github.com/lucasvmiguel/go-analytics/controllers/websocket"
)

func Start(debugger bool, recovery bool, version string, port string) {
	r := gin.New()
	
	if debugger {
		r.Use(gin.Logger())
	}

	if recovery {
		r.Use(gin.Recovery())
	}

	r.Use(corsMiddleware())

	r.POST("/"+version+"/auth", controllers.AccessController)
	r.Use(controllers.AuthentificationController)
	r.POST("/"+version+"/notification", controllers.NotificationController)
	r.GET("/"+version+"/ws", websocket.WebsocketController)

	err := r.Run(port)
	if err != nil {
		logrus.Panic("Error to open api")
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
