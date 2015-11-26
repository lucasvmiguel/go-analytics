package standard

import (
	"github.com/Sirupsen/logrus"

	"gopkg.in/mgo.v2"
)

const NAME_TABLE = "companies"

var collection *mgo.Collection

func Start(addr string, dbname string, username string, password string) {
	session, err := mgo.Dial(addr)
	if err != nil {
		logrus.Panic("connect to database standard fail")
	}

	collection = session.DB(dbname).C(NAME_TABLE)
	logrus.Info("connected to standard database!")
}
