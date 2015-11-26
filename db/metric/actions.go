package metric

import (
	"time"

	"github.com/Sirupsen/logrus"
	database "github.com/influxdb/influxdb/client/v2"
	"github.com/lucasvmiguel/go-analytics/model"
)

const NAME_TABLE = "notifications"

func SaveNotification(noti model.Notification) {
	if clientDB.BP == nil || clientDB.Client == nil {
		logrus.Error("something is wrong with database client")
		return
	}

	pt, err := database.NewPoint(NAME_TABLE, nil, noti.ToMapString(), time.Now())
	if err != nil {
		logrus.Error("New notification error")
		return
	}

	clientDB.BP.AddPoint(pt)
	err = clientDB.Client.Write(clientDB.BP)
	if err != nil {
		logrus.Error("error to write new notification")
	}
	logrus.Info("notification writed")
}
