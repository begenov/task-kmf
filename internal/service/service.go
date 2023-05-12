package service

import (
	"context"

	"github.com/begenov/tesk-kmf/internal/model"
	"github.com/begenov/tesk-kmf/internal/repository"
	currencyservice "github.com/begenov/tesk-kmf/internal/service/currency_service"
)

type CurrencyIR interface {
	CreateCurrency(ctx context.Context, rates model.Rates) error
}

type Service struct {
	Currency CurrencyIR
}

func NewService(bankDB repository.BankDB) *Service {
	return &Service{
		Currency: currencyservice.NewCurrencyService(bankDB),
	}
}
