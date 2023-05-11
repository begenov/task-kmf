package model

import "time"

type Currency struct {
	ID    int       `json:"id"`
	Title string    `json:"Title"`
	Code  string    `json:"code"`
	Value float64   `json:"value"`
	ADate time.Time `json:"ADATE"`
}
