package main

import (
	"log"

	"github.com/begenov/tesk-kmf/internal/app"
	"github.com/begenov/tesk-kmf/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	app := app.NewApplication(cfg)

	if err := app.Run(); err != nil {
		return
	}
}
