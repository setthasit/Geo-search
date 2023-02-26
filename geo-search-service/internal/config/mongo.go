package config

import (
	"geo-search/internal/utils"

	"github.com/Netflix/go-env"
)

type MongoConfig struct {
	URI      string `env:"APP_MONGO_URI,required=true"`
	Username string `env:"MONGO_USERNAME,required=true"`
	Password string `env:"MONGO_PASSWORD,required=true"`
	Database string `env:"MONGO_DB,required=true"`
}

func WireMongoConfig(l ConfigLoader) *MongoConfig {
	err := l.Load()
	if err != nil {
		utils.LogWarn(err.Error())
	}

	cfg := &MongoConfig{}
	_, err = env.UnmarshalFromEnviron(cfg)
	if err != nil {
		utils.LogFatal(err.Error())
	}

	return cfg
}
