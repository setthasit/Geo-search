package config

import (
	"flag"
	"os"
)

type CommandConfig struct {
	EnvFile string
}

func WireCommandConfig() *CommandConfig {
	cfg := &CommandConfig{}

	if os.Getenv("IT_TEST") != "true" {
		file := flag.String("env_file", ".env", "specific file for to load ENV values")
		flag.Parse()
		cfg.EnvFile = *file
	} else {
		cfg.EnvFile = "../../.env.test"
	}

	return cfg
}
