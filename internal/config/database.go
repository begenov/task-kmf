package config

type database struct {
	Driver string `json:"driver"`
	DSN    string `json:"dsn"`
}
