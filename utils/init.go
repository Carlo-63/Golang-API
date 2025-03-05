package utils

import (
	"encoding/json"
	"main/models"
	"os"
)

var config models.Config

func Init() models.Config {
	data, err := os.ReadFile("config/config.json")

	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	return config
}
