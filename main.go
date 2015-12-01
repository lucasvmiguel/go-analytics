package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/lucasvmiguel/go-analytics/api"

	db_metric "github.com/lucasvmiguel/go-analytics/db/metric"

	"github.com/spf13/viper"
)

func init() {
	path, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		panic("Can't read config file")
	}
}

func main() {
	environment := os.Args[1]
	if environment == "" {
		logrus.Panic("You need set the enviroment")
	}

	db_metric.Start(
		viper.GetString(environment+".db.metric.addr"),
		viper.GetString(environment+".db.metric.dbname"),
		viper.GetString(environment+".db.metric.username"),
		viper.GetString(environment+".db.metric.password"),
	)

	api.Start(
		viper.GetBool(environment+".debugger"),
		viper.GetBool(environment+".recovery"),
		viper.GetString(environment+".api.version"),
		viper.GetString(environment+".api.port"),
	)
}
