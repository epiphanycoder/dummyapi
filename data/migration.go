package data

import (
	"dummyapi/db"
	"dummyapi/model"
	"github.com/sirupsen/logrus"
)

type Migration struct {
	adb      *db.AuthUserDb
	dataFile string
}

const passowrd = "pass1234"

func NewMigration(adb *db.AuthUserDb, file string) *Migration {
	return &Migration{adb: adb, dataFile: file}
}

func (m *Migration) Run() {
	logrus.Debugf("Data migration start")

	logrus.Debugf("Data migration end")
}

func (m *Migration) createUsers(users *[]*model.User) {

}
