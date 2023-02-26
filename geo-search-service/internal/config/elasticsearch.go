package config

import (
	"geo-search/internal/utils"

	"github.com/Netflix/go-env"
)

type ElasticsearchConfig struct {
	URI      string `env:"APP_ELASTICSEARCH_URI"`
	Username string `env:"APP_ELASTICSEARCH_USERNAME"`
	Password string `env:"APP_ELASTICSEARCH_PASSWORD"`
	CertPath string `env:"APP_ELASTICSEARCH_CERT_PATH"`
}

func WireElasticsearchConfig(l ConfigLoader) *ElasticsearchConfig {
	err := l.Load()
	if err != nil {
		utils.LogWarn(err.Error())
	}

	cfg := &ElasticsearchConfig{}
	_, err = env.UnmarshalFromEnviron(cfg)
	if err != nil {
		utils.LogFatal(err.Error())
	}

	return cfg
}
