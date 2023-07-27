package config

import (
	"strings"
	"sync"

	"github.com/JhonatanRSantos/gorest/internal/platform/goenv"
)

var (
	defaultAppEnv  = string(goenv.Local)
	defaultAppName = "ms-backend-gorest"
	defaultAppPort = "8080"
)

var (
	config         *Config
	initConfigOnce sync.Once
)

type Config struct {
	AppEnv  goenv.Env
	AppName string
	AppPort string
}

// GetConfig Get the environment vars
func GetConfig() *Config {
	initConfigOnce.Do(func() {
		if config == nil {
			config = &Config{}
			config.AppEnv = goenv.Env(strings.ToLower(goenv.Load("APP_ENV", defaultAppEnv)))
			config.AppName = goenv.Load("APP_NAME", defaultAppName)
			config.AppPort = goenv.Load("APP_PORT", defaultAppPort)
		}
	})
	return config
}
