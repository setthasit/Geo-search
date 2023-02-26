package config

import (
	"geo-search/internal/utils"

	"github.com/Netflix/go-env"
)

type EnvironmentType string

const (
	Development EnvironmentType = "development"
	Production  EnvironmentType = "production"
)

func (et EnvironmentType) String() string {
	return string(et)
}

func (et EnvironmentType) Valid() bool {
	switch et {
	case Development:
		return true
	case Production:
		return true
	default:
		return false
	}
}

type AppConfig struct {
	Environment EnvironmentType `env:"APP_ENVIRONMENT,default=production"`
}

func WireAppConfig(l ConfigLoader) *AppConfig {
	err := l.Load()
	if err != nil {
		utils.LogWarn(err.Error())
	}

	appConfig := &AppConfig{}
	_, err = env.UnmarshalFromEnviron(appConfig)
	if err != nil {
		utils.LogFatal(err.Error())
	}

	if !appConfig.Environment.Valid() {
		utils.LogWarn("invalid APP_ENVIRONMENT fallback to production")
		appConfig.Environment = Production
	}

	return appConfig
}
