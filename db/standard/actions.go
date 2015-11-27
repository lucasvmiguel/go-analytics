package standard

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/lucasvmiguel/go-analytics/model"
)

func MockCompany() {
	err := client.Set("lucas", "123456", 0)
	if err != nil {
		fmt.Println(err)
	}
}

func GetCompanyName(key string) model.Company {

	val, err := client.Get(key).Result()
	if err != nil {
		logrus.Error("can't find company")
	}

	return model.Company{}
}
