package standard

import (
	"github.com/Sirupsen/logrus"

	"gopkg.in/redis.v3"
)

const NAME_TABLE = "companies"

var client *redis.Client

func Start(addr string, dbname string, username string, password string) {

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		logrus.Panic("connect to database standard fail")
	}
}
