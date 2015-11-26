package controllers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/lucasvmiguel/go-analytics/db/metric"
	"github.com/lucasvmiguel/go-analytics/model"
)

func NotificationController(c *gin.Context) {

	notification := model.Notification{}

	token := c.Request.Header.Get("authorization")
	body, _ := ioutil.ReadAll(c.Request.Body)

	err := json.Unmarshal([]byte(body), &notification)
	if err != nil || token == "" {
		logrus.Info("error to serialize notification")
		c.AbortWithStatus(422)
		return
	}

	go saveRoutine(token, notification)
	c.AbortWithStatus(200)
}

func saveRoutine(token string, noti model.Notification) {
	metric.SaveNotification(noti)
}
