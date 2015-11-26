package controllers

import "github.com/gin-gonic/gin"

func AccessController(c *gin.Context) {
	c.String(200, "access controller")
}

func AuthentificationController(c *gin.Context) {
	//c.String(403, "authentification controller")
	c.Next()
}
