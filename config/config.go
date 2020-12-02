package config

import (
	"github.com/jason-shen/gopush/pkg/utils/logger"
	"gopkg.in/ini.v1"
	"os"
)

type Config struct {
	Database 	*DatabaseConfig
	Jwt			*JwtConfig
}

func New() *Config {
	return &Config{
		Database: 	NewDatabase(),
		Jwt: 		NewJwtConfig(),
	}
}

func getIni(section, key, defaultValue string) string {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		logger.Errorf("Fail to read file: %v", err)
		os.Exit(1)
	}

	if value := cfg.Section(section).Key(key).String(); value != "" {
		return value
	}
	return defaultValue
}
