package main

import (
	"os"

	"github.com/lucasvmiguel/go-analytics/api"
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
		panic("You need set the enviroment")
	}

	api.Start(viper.GetString(environment+".api.version"), viper.GetString(environment+".api.port"))
}
