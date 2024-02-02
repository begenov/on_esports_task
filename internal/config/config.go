package config

import (
	"on_esports/internal/logger"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	OPEN_WEATHER_MAP_API_KEY         string `envconfig:"OPEN_WEATHER_MAP_API_KEY"`
	GOOGLE_CLOUD_TRANSLATION_API_KEY string `envconfig:"GOOGLE_CLOUD_TRANSLATION_API_KEY"`
	DISCORD_BOT_TOKEN                string `envconfig:"DISCORD_BOT_TOKEN"`
}

var Config config

func New(fpath string) error {

	err := godotenv.Load(fpath)
	if err != nil {
		logger.Errorf("godotenv.Load(): %v", err.Error())
		return err

	}
	err = envconfig.Process("", &Config)
	if err != nil {
		logger.Errorf("envconfig.Process(): %v", err.Error())
	}

	return err
}
