package scraper

import (
	"errors"
	"fmt"
	"hotelboss/app/scraper/hotels/marriot"
	"hotelboss/app/scraper/hotels/parkroyal"
	"hotelboss/app/scraper/types"
	"hotelboss/internal/client/ssllabs"
	"hotelboss/internal/client/whois"
)

var scrapers = map[string]types.ScrapeHotelFunc{
	marriot.URL:   marriot.Scrape,
	parkroyal.URL: parkroyal.Scrape,
}

type Scraper interface {
	ScrapeHotel(url string) (*types.HotelInfo, error)
}

type scraper struct{}

func New() Scraper {
	return &scraper{}
}

// ScrapeHotel implements Scraper.
func (*scraper) ScrapeHotel(url string) (*types.HotelInfo, error) {
	scrap, ok := scrapers[url]
	if !ok {
		return nil, errors.New(fmt.Sprintf("scraper for site %s not found", url))
	}
	site, err := scrap(url)
	if err != nil {
		return nil, err
	}

	hostInfo, err := ssllabs.Analyze(url)
	if err != nil {
		return nil, err
	}
	info := &types.HostInfo{
		Protocol: hostInfo.Protocol,
	}
	for _, ept := range hostInfo.Endpoints {
		info.Endpoints = append(info.Endpoints, types.Endpoint{Name: ept.ServerName})
	}

	wis, err := whois.Whois(url)
	if err != nil {
		return nil, err
	}
	info.CreationDate = wis.CreationDate
	info.ExpirationDate = wis.ExpirationDate
	info.Owner = wis.Owner
	info.ContactEmail = wis.ContactEmail

	return &types.HotelInfo{Site: site, Host: info}, nil
}
