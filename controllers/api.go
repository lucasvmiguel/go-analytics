package controllers

import "github.com/gin-gonic/gin"

func NotificationController(c *gin.Context) {
	c.String(200, "notification controller")
}
