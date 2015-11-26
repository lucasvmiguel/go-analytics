package standard

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/Sirupsen/logrus"

	"github.com/lucasvmiguel/go-analytics/model"
)

func MockCompany() {
	err := collection.Insert(&model.Company{1, "Lucas", "123456"})
	if err != nil {
		fmt.Println(err)
	}
}

func GetCompany(id uint) model.Company {
	result := model.Company{}

	err := collection.Find(bson.M{"ID": id}).One(&result)
	if err != nil {
		logrus.Error("can't find company")
	}

	return result
}
