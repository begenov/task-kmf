package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Database database `json:"database"`
	Server   server   `json:"server"`
}

const path = "./internal/config/config.json"

func NewConfig() (*Config, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config

	if err = json.Unmarshal(body, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
