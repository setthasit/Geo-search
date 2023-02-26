package config

import (
	"github.com/joho/godotenv"
)

//go:generate mockgen -source loader.go -destination mock_config/mock_loader.go -package mock_config
type ConfigLoader interface {
	Load() error
}

type ConfigLoaderImpl struct {
	loaded bool

	commandConfig *CommandConfig
}

func WireConfigLoader(commandConfig *CommandConfig) *ConfigLoaderImpl {
	return &ConfigLoaderImpl{
		commandConfig: commandConfig,
	}
}

func (cfg *ConfigLoaderImpl) Load() error {
	if cfg.loaded { // Check if already loaded ENV vars
		return nil // if its loaded then skip
	}

	err := godotenv.Load(cfg.commandConfig.EnvFile)
	if err != nil {
		cfg.loaded = true
		return err
	}

	cfg.loaded = true

	return nil
}
