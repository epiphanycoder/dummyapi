package cfg

import (
	"dummyapi/util"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"strings"
)

const baseConfigFilePath = "dummyapi.json"

//Config contains configuration data
type Config struct {
	configData *configData
	logDebug   bool
	appName    string
	version    string
}

type configData struct {
	AppName        string
	Version        string
	Description    string
	Server         ServerDef
	Indent         bool
	RouteConfig    RouteConfig
	JwtTokenConfig JwtTokenConfig
	AuthUserDb     DbDef
	DummyApiConfig DummyApiDef
	UserDataPath   string
	configFilePath string
	Logging        logDef
}

type ServerDef struct {
	Bind string
	Port int
}

type RouteConfig struct {
	StrictSlash bool
	SkipClean   bool
}

// DummyApiDef defines dummyapi.io definition
type DummyApiDef struct {
	AppId      string
	BaseURL    string
	UserAgent  string
	PageConfig DummyApiPageConfig
}

type DummyApiPageConfig struct {
	DefaultPageSize int
}

//JwtTokenConfig defines jwt config details
type JwtTokenConfig struct {
	Secret    string
	ExpiresAt int
}

// logDef defines logging
type logDef struct {
	Filename string
	Level    string
}

func (c *Config) AppName() string {
	return c.appName
}

func (c *Config) Version() string {
	return c.version
}

func (c *Config) ServerDef() *ServerDef {
	return &c.configData.Server
}

func (c *Config) AuthUserDbDef() *DbDef {
	return &c.configData.AuthUserDb
}

func (c *Config) DummyApiDef() *DummyApiDef {
	return &c.configData.DummyApiConfig
}

func (ld *logDef) isDebug() bool {
	return strings.EqualFold(ld.Level, "DEBUG")
}

// NewConfig creates the app configuration
func NewConfg() *Config {
	cd := loadConfig()
	ld := cd.Logging.isDebug()
	return &Config{cd, ld, cd.AppName, cd.Version}
}

func loadConfig() *configData {
	flag.Parse()
	configFilePath := flag.Arg(0)

	if util.IsEmpty(configFilePath) {
		configFilePath = baseConfigFilePath
	}
	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", configFilePath, err)
	}
	var configData configData
	if err := json.Unmarshal(data, &configData); err != nil {
		log.Fatalf("failed to unmarshal file %s : %v", configFilePath, err)
	}
	configData.configFilePath = configFilePath
	return &configData
}
