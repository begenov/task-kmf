package model

type Currency struct {
	ID          int
	FullName    string `xml:"fullname"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	ADate       string
}
