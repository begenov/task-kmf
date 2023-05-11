package currencyservice

import (
	"context"

	"github.com/begenov/tesk-kmf/internal/model"
)

type CurrencyService struct {
	povider bankProvider
}

type bankProvider interface {
}

func NewCurrencyService(proivder bankProvider) *CurrencyService {
	return &CurrencyService{
		povider: proivder,
	}
}

func (service *CurrencyService) CreateCurrency(ctx context.Context, currency model.Currency) error {
	return nil
}
