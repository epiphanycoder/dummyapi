package cfg

import (
	"dummyapi/util"
	"os"
	"strconv"
)

//DbDef defnes database definition
type DbDef struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	Logging  bool
}

// DbUser returns configured db user or the user defined in the environment variable's value
func (dbd *DbDef) DbUser() string {
	if util.IsEmpty(os.Getenv("DB_USER")) {
		return dbd.User
	}
	return os.Getenv("DB_USER")
}

// DbPass returns configured db user's password or environment variable's value
func (dbd *DbDef) DbPass() string {
	if util.IsEmpty(os.Getenv("DB_PASSWORD")) {
		return dbd.Password
	}
	return os.Getenv("DB_PASSWORD")
}

// DbHost returns configured db host or environment variable's value
func (dbd *DbDef) DbHost() string {
	if util.IsEmpty(os.Getenv("DB_HOST")) {
		return dbd.Host
	}
	return os.Getenv("DB_HOST")
}

// DbPort returns configured db port or environment variable's value
func (dbd *DbDef) DbPort() int {
	if util.IsEmpty(os.Getenv("DB_PORT")) {
		return dbd.Port
	}
	if port, err := strconv.Atoi(os.Getenv("DB_PORT")); err != nil {
		return dbd.Port
	} else {
		return port
	}
}

func (dbd *DbDef) DbName() string {
	if util.IsEmpty(os.Getenv("DB_NAME")) {
		return dbd.Database
	}
	return os.Getenv("DB_NAME")
}
