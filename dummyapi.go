package dummyapi

import (
	"dummyapi/cfg"
	"dummyapi/data"
	"dummyapi/db"
	"dummyapi/logger"
)

func main() {

	c := cfg.NewConfg()
	lf := logger.Configure(c.LogFile(), c.LogLevel())
	defer lf.Close()

	adb := db.NewAuthUserDb(c.AuthUserDbDef())
	mig := data.NewMigration(adb, c.UserDataPath())
	mig.Run()
}
