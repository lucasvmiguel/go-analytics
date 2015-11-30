package standard

import (
	"fmt"

	"github.com/lucasvmiguel/go-analytics/model"
	"github.com/spf13/viper"
)

func MockCompany() {
	err := client.Set("lucas", "123456", 0)
	if err != nil {
		fmt.Println(err)
	}
}

func GetCompanyName(key string) model.Company {
	company := model.Company{}

	c := viper.GetStringMap("companies")[key]

	if c != nil {
		company.Name = c.(map[string]interface{})["name"].(string)
		company.Key = key
	}
	return company
}
