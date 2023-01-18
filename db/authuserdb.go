package db

import (
	"dummyapi/cfg"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

//AuthUserDb defines auth user db
type AuthUserDb struct {
	Db *gorm.DB
}

func NewAuthUserDb(dbDef *cfg.DbDef) *AuthUserDb {
	return &AuthUserDb{connectDatabase(dbDef)}
}

func connectDatabase(dbDef *cfg.DbDef) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%d:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbDef.DbUser(), dbDef.DbPass(), dbDef.DbHost(), dbDef.DbPort(), dbDef.DbName())
	var gormCfg *gorm.Config

	if dbDef.Logging {
		gormCfg = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	}

	db, err := gorm.Open(mysql.Open(dsn), gormCfg)
	if err != nil {
		log.Fatalf("failed to connect database %s:%d", dbDef.DbHost(), dbDef.DbPort())
	}
	return db
}
