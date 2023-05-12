package currencyservice

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/begenov/tesk-kmf/internal/model"
)

type CurrencyService struct {
	povider bankProvider
}

type bankProvider interface {
	CreateCurrency(ctx context.Context, rates model.Rates) error
	CurrencyByCode(ctx context.Context, date time.Time, code string) ([]model.Currency, error)
	CurrencyByDate(ctx context.Context, date time.Time) ([]model.Currency, error)
}

func NewCurrencyService(proivder bankProvider) *CurrencyService {
	return &CurrencyService{
		povider: proivder,
	}
}

func (service *CurrencyService) CreateCurrency(ctx context.Context, rates model.Rates, caster chan error) {
	var err error
	rates.DataTime, err = formatDate(rates.Data)
	if err != nil {
		log.Println(err)
		caster <- err
		return
	}

	err = service.povider.CreateCurrency(ctx, rates)
	if err != nil {
		fmt.Printf("Error while saving currency rates to DB: %v", err)
		caster <- err
		return

	}
	fmt.Printf("Data for %s saved\n", rates.Data)
	caster <- nil
}

func (service *CurrencyService) GetCurrencyByCode(ctx context.Context, date string, code string) ([]model.Currency, error) {
	dataTime, err := formatDate(date)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return service.povider.CurrencyByCode(ctx, dataTime, code)
}

func (service *CurrencyService) GetCurrency(ctx context.Context, data string) ([]model.Currency, error) {
	dataTime, err := formatDate(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return service.povider.CurrencyByDate(ctx, dataTime)
}

func formatDate(d string) (time.Time, error) {
	format := "02.01.2006"
	date, err := time.Parse(format, d)
	if err != nil {
		fmt.Println(err)
		return time.Time{}, err
	}
	return date, nil
}
