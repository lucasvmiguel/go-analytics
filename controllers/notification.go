package controllers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/lucasvmiguel/go-analytics/controllers/websocket"
	"github.com/lucasvmiguel/go-analytics/db/metric"
	"github.com/lucasvmiguel/go-analytics/db/standard"
	"github.com/lucasvmiguel/go-analytics/errors"
	"github.com/lucasvmiguel/go-analytics/model"
	"github.com/spf13/viper"
)

func NotificationController(c *gin.Context) {

	notification := model.Notification{}

	body, _ := ioutil.ReadAll(c.Request.Body)

	err := json.Unmarshal([]byte(body), &notification)
	if err != nil {
		logrus.Error(errors.ErrSerializeNotification.Error())
		c.AbortWithStatus(422)
		return
	}

	go saveRoutine(notification)
	c.AbortWithStatus(200)
}

func saveRoutine(noti model.Notification) {
	noti.Company = standard.GetCompanyName(noti.Company).Name

	if noti.Company == "" {
		logrus.Error(errors.ErrFindCompany.Error())
		return
	}

	go metric.SaveNotification(noti)
	needSend(noti)
}

func needSend(noti model.Notification) {
	if noti.Type >= uint8(viper.GetInt("develop.websocket.condition.type")) &&
		noti.Relevance >= uint8(viper.GetInt("develop.websocket.condition.relevance")) {
		go websocket.Send(noti)
	}
}
