package controllers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/lucasvmiguel/go-analytics/controllers/websocket"
	"github.com/lucasvmiguel/go-analytics/db/metric"
	"github.com/lucasvmiguel/go-analytics/db/standard"
	"github.com/lucasvmiguel/go-analytics/model"
)

func NotificationController(c *gin.Context) {

	notification := model.Notification{}

	body, _ := ioutil.ReadAll(c.Request.Body)

	err := json.Unmarshal([]byte(body), &notification)
	if err != nil {
		logrus.Error("error to serialize notification")
		c.AbortWithStatus(422)
		return
	}

	go saveRoutine(notification)
	c.AbortWithStatus(200)
}

func saveRoutine(noti model.Notification) {
	noti.Company = standard.GetCompanyName(noti.Company).Name

	if noti.Company == "" {
		logrus.Error("error to find company")
		return
	}

	go metric.SaveNotification(noti)

	if noti.Type == model.ERROR && noti.Relevance == model.HIGH {
		go websocket.Send(noti)
	}
}
