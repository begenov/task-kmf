package controller

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"

	"log"
	"net/http"

	"github.com/begenov/tesk-kmf/internal/model"
	"github.com/gorilla/mux"
)

const api = "https://nationalbank.kz/rss/get_rates.cfm?fdate=%s"

func (c *Controller) saveCurrencyDate(w http.ResponseWriter, r *http.Request) {
	date := mux.Vars(r)["date"]

	resp, err := http.Get(fmt.Sprintf(api, date))

	if err != nil {
		log.Printf("Error when making a request to the national bank's API: %v", err)
		http.Error(w, "Error when making a request to the national bank's API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}
	var rates model.Rates

	if err = xml.Unmarshal(body, &rates); err != nil {
		log.Printf("Error parsing response body: %v", err)
		http.Error(w, "Error parsing response body", http.StatusInternalServerError)
		return
	}

	go c.service.Currency.CreateCurrency(context.Background(), rates)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"success": true})

}

func (c *Controller) currencyHandler(w http.ResponseWriter, r *http.Request) {
	data := mux.Vars(r)["date"]
	code := mux.Vars(r)["code"]

	var currency []model.Currency
	var err error

	if code != "" {
		currency, err = c.servicecCurrency.GecCurrencyByCode(context.Background(), data, code)
	} else {
		currency, err = c.servicecCurrency.GetCurrency(context.Background(), data)
	}

	if err != nil {
		log.Printf("Error getting currency: %v", err)
		http.Error(w, "Error getting currency", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(currency)

}
