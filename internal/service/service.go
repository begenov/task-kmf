package service

import (
	"context"

	"github.com/begenov/tesk-kmf/internal/model"
	"github.com/begenov/tesk-kmf/internal/repository"
)

type CurrencyIR interface {
	CreateCurrency(ctx context.Context, currency model.Currency) error
}

type Service struct {
	Currency CurrencyIR
}

func NewService(bankDB repository.BankDB) *Service {
	return &Service{bankDB}
}
