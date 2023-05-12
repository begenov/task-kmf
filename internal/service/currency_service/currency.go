package currencyservice

import (
	"context"
	"fmt"
	"time"

	"github.com/begenov/tesk-kmf/internal/model"
)

type CurrencyService struct {
	povider bankProvider
}

type bankProvider interface {
	CreateCurrency(ctx context.Context, rates model.Rates) error
}

func NewCurrencyService(proivder bankProvider) *CurrencyService {
	return &CurrencyService{
		povider: proivder,
	}
}

func (service *CurrencyService) CreateCurrency(ctx context.Context, rates model.Rates) error {
	format := "02.01.2006"
	date, err := time.Parse(format, rates.Data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	rates.DataTime = date
	err = service.povider.CreateCurrency(ctx, rates)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("Data for %s saved\n", rates.Data)
	return nil
}

func Currency(ctx context.Context, date string, code string) ([]model.Currency, error) {
}
