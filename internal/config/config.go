package config

import (
	"encoding/json"
	"log"
	"os"
)

type appMode string

const (
	MODE_DEVELOPMENT appMode = "development"
	MODE_RELEASE     appMode = "release"
)

type appConfig struct {
	Name  string  `json:"app_name"`
	Mode  appMode `json:"app_mode"`
}

func NewAppConfig(filePath string) appConfig {
	config := appConfig{}

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	switch config.Mode {
	case MODE_DEVELOPMENT:
		log.Println("The App is running in development mode")
	case MODE_RELEASE:
		log.Println("The App is running in release mode")
	default:
		log.Fatalf("Unknow App mode: %s\n", config.Mode)
	}

	return config
}
