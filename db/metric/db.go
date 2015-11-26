package metric

import (
	"github.com/Sirupsen/logrus"
	database "github.com/influxdb/influxdb/client/v2"
)

type DBClient struct {
	Client database.Client
	BP     database.BatchPoints
}

var clientDB DBClient

func Start(addr string, dbname string, username string, password string) {
	var err error

	clientDB.Client, err = database.NewHTTPClient(database.HTTPConfig{
		Addr:     addr,
		Username: username,
		Password: password,
	})

	if err != nil {
		logrus.Panic("connect to metric database fail")
	}

	clientDB.BP, err = database.NewBatchPoints(database.BatchPointsConfig{
		Database:  dbname,
		Precision: "s",
	})

	logrus.Info("connected to metric database!")
}
