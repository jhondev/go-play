package models

type Location struct {
	City    string `json:"city"`
	Country string `json:"country"`
	Address string `json:"address"`
	ZipCode string `json:"zip_code"`
}
