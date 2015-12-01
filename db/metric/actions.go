package metric

import (
	"time"

	"github.com/Sirupsen/logrus"
	database "github.com/influxdb/influxdb/client/v2"
	"github.com/lucasvmiguel/go-analytics/errors"
	"github.com/lucasvmiguel/go-analytics/model"
)

const NAME_TABLE = "notifications"

func SaveNotification(noti model.Notification) {
	if clientDB.BP == nil || clientDB.Client == nil {
		logrus.Error(errors.ErrInvalidConnectionDatabase.Error())
		return
	}

	pt, err := database.NewPoint(NAME_TABLE, nil, noti.ToMapString(), time.Now())
	if verifyError(errors.ErrNewNotification.Error(), err) {
		return
	}

	clientDB.BP.AddPoint(pt)
	err = clientDB.Client.Write(clientDB.BP)
	if verifyError(errors.ErrNewNotification.Error(), err) {
		return
	}

	logrus.Info("notification writed")
}

func verifyError(err string, errs ...error) bool {
	for _, e := range errs {
		if e != nil {
			logrus.Error(err)
			return true
		}
	}
	return false
}
