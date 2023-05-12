package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"

	"log"
	"net/http"

	"github.com/begenov/tesk-kmf/internal/model"
	"github.com/gorilla/mux"
)

const api = "https://nationalbank.kz/rss/get_rates.cfm?fdate=%s"

var (
	success = map[string]bool{"success": true}
)

type fail struct {
	Error bool
	Msg   string
}

func (c *Controller) saveCurrencyDate(w http.ResponseWriter, r *http.Request) {
	date := mux.Vars(r)["date"]

	resp, err := http.Get(fmt.Sprintf(api, date))

	if err != nil {
		log.Printf("Error when making a request to the national bank's API: %v", err)
		jsonEncoding(w, http.StatusBadRequest, fail{
			Error: false,
			Msg:   "error request",
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		jsonEncoding(w, http.StatusBadRequest, fail{
			Error: false,
			Msg:   "error request",
		})
		return
	}
	var rates model.Rates

	if err = xml.Unmarshal(body, &rates); err != nil {
		log.Printf("Error parsing response body: %v", err)
		jsonEncoding(w, http.StatusBadRequest, fail{
			Error: false,
			Msg:   "error request",
		})
		return
	}
	caster := make(chan error)

	go c.service.Currency.CreateCurrency(context.Background(), rates, caster)
	err = <-caster
	if err != nil {
		jsonEncoding(w, http.StatusBadRequest, fail{
			Error: false,
			Msg:   "error request",
		})
		return
	}
	jsonEncoding(w, http.StatusOK, success)
}

func (c *Controller) currencyHandler(w http.ResponseWriter, r *http.Request) {
	data := mux.Vars(r)["date"]
	code := mux.Vars(r)["code"]

	var currency []model.Currency
	var err error

	if code != "" {
		currency, err = c.service.Currency.GetCurrencyByCode(context.Background(), data, code)
	} else {
		fmt.Println(code, "code")
		currency, err = c.service.Currency.GetCurrency(context.Background(), data)
	}

	if err != nil {
		log.Printf("Error getting currency: %v", err)
		if errors.Is(err, sql.ErrNoRows) {
			jsonEncoding(w, http.StatusNotFound, fail{
				Error: false,
				Msg:   "Data not found for the given data and code",
			})
			return
		}
		jsonEncoding(w, http.StatusBadRequest, fail{
			Error: false,
			Msg:   "Invalid date format",
		})
		return
	}
	if len(currency) == 0 {
		jsonEncoding(w, http.StatusNotFound, fail{
			Error: false,
			Msg:   "Data not found for the given data and code",
		})
		return
	}
	jsonEncoding(w, http.StatusOK, currency)
}

func jsonEncoding(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
