package config

import (
	"fmt"
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
	fmt.Println(body)
	return &Config{}, nil
}
