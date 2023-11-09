package core

import (
	cstore "hotelboss/app/company/store"
	"hotelboss/app/franchise/store"
	"hotelboss/app/scraper"
)

type Franchise struct {
	scraper.Scraper
	store        store.FranchiseStore
	companyStore cstore.CompanyStore
}

func New(s scraper.Scraper, st store.FranchiseStore, cst cstore.CompanyStore) *Franchise {
	return &Franchise{Scraper: s, store: st, companyStore: cst}
}
