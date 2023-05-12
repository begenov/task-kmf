package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (c *Controller) InitRouter() *mux.Router {
	mux := mux.NewRouter()

	mux.HandleFunc("/currency/save/{date}", c.saveCurrencyDate).Methods(http.MethodGet)
	mux.HandleFunc("/currency/{date}", c.currencyHandler).Methods(http.MethodGet)
	mux.HandleFunc("/currency/{date}/{code}", c.currencyHandler).Methods(http.MethodGet)

	return mux
}
