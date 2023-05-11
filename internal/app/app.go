package app

import "github.com/begenov/tesk-kmf/internal/config"

type Application struct {
	cfg *config.Config
}

func NewApplication(cfg *config.Config) *Application {
	return &Application{
		cfg: cfg,
	}
}

func (app *Application) Run() error {
	return nil
}
