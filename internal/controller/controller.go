package controller

import "github.com/begenov/tesk-kmf/internal/service"

type Controller struct {
	service *service.Service
}

func NewContoller(service *service.Service) *Controller {
	return &Controller{
		service: service,
	}
}
