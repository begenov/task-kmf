package main

import (
	"github.com/begenov/tesk-kmf/internal/app"
	"github.com/begenov/tesk-kmf/internal/config"
)

func main() {
	cfg := config.NewConfig()
	app := app.NewApplication(cfg)

	if err := app.Run(); err != nil {
		return
	}
}
