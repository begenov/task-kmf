package model

import "time"

type Rates struct {
	Data     string     `xml:"date"`
	Currency []Currency `xml:"item"`
	DataTime time.Time
}
