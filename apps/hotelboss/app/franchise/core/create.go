package core

import (
	"fmt"
	cmodels "hotelboss/app/company/models"
	fmodels "hotelboss/app/franchise/models"

	"github.com/gofrs/uuid"
)

func (f *Franchise) Create(c *cmodels.Company, fchs []fmodels.FranchiseDTO) error {
	err := f.companyStore.Create(c)
	if err != nil {
		return err
	}

	for _, fch := range fchs {
		fmt.Println("scraping hotel", fch.URL)
		hinfo, err := f.ScrapeHotel(fch.URL)
		if err != nil {
			return err
		}
		id, _ := uuid.NewV7()
		f.store.Create(&fmodels.NewFranchise{
			ID:         id.String(),
			Name:       fch.Name,
			URL:        fch.URL,
			LogoURL:    hinfo.Site.LogoURL,
			LocationID: fch.Location.ZipCode, // TODO: fix locationID
		})
	}

	return nil
}
