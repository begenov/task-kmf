package main

import (
	"log"

	"github.com/begenov/tesk-kmf/internal/app"
	"github.com/begenov/tesk-kmf/internal/config"
)

// @title TASK-KMF
// @version 1.0
// @description This is test task

// @host localhost:4000
// @BasePath /

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
