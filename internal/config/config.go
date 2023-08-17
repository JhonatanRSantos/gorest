package config

import (
	"strings"
	"sync"

	"github.com/JhonatanRSantos/gorest/internal/platform/goenv"
)

var (
	defaultAppEnv    = string(goenv.Local)
	defaultAppName   = "ms-backend-gorest"
	defaultAppPort   = "8080"
	defaultAppDBHost = "127.0.0.1"
	defaultAppDBPort = int64(5432)
	defaultAppDBUser = ""
	defaultAppDBPass = ""
	defaultAppDBName = ""
)

var (
	config         *Config
	initConfigOnce sync.Once
)

type DBConfig struct {
	Host   string
	Port   int64
	User   string
	Pass   string
	DBName string
}
type Config struct {
	AppEnv   goenv.Env
	AppName  string
	AppPort  string
	DBConfig DBConfig
}

// GetConfig Get the environment vars
func GetConfig() *Config {
	initConfigOnce.Do(func() {
		if config == nil {
			config = &Config{}
			config.AppEnv = goenv.Env(strings.ToLower(goenv.Load("APP_ENV", defaultAppEnv)))
			config.AppName = goenv.Load("APP_NAME", defaultAppName)
			config.AppPort = goenv.Load("APP_PORT", defaultAppPort)
			config.DBConfig.Host = goenv.Load("APP_DB_HOST", defaultAppDBHost)
			config.DBConfig.Port = goenv.Load("APP_DB_PORT", defaultAppDBPort)
			config.DBConfig.User = goenv.Load("APP_DB_USER", defaultAppDBUser)
			config.DBConfig.Pass = goenv.Load("APP_DB_PASS", defaultAppDBPass)
			config.DBConfig.DBName = goenv.Load("APP_DB_NAME", defaultAppDBName)
		}
	})
	return config
}
