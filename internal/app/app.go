package app

import (
	"log"
	"net/http"

	"github.com/begenov/tesk-kmf/internal/config"
	"github.com/begenov/tesk-kmf/internal/controller"
	"github.com/begenov/tesk-kmf/internal/repository"
	"github.com/begenov/tesk-kmf/internal/service"
	"github.com/begenov/tesk-kmf/pkg/db"
)

type Application struct {
	cfg *config.Config
}

func NewApplication(cfg *config.Config) *Application {
	return &Application{
		cfg: cfg,
	}
}

func (app *Application) Run() error {
	db, err := db.OpenDB(app.cfg.Database.Driver, app.cfg.Database.DSN)
	if err != nil {
		return err
	}

	repository := repository.NewRepository(db)

	service := service.NewService(repository.BankDB)

	controller := controller.NewContoller(service)

	log.Printf("Starting server in port = %s\n", app.cfg.Server.Port)

	return http.ListenAndServe(app.cfg.Server.Port, controller.InitRouter())
}
