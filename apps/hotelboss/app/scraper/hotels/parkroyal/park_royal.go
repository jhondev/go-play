package parkroyal

import (
	"hotelboss/app/scraper/types"

	"github.com/gocolly/colly/v2"
)

const URL = "www.park-royalhotels.com"

func Scrape(url string) (*types.HotelSiteInfo, error) {
	var e error
	info := &types.HotelSiteInfo{}

	c := colly.NewCollector()
	// Find and get logo link
	c.OnHTML("a.logo > img", func(h *colly.HTMLElement) {
		info.LogoURL = h.Attr("src")
	})
	c.OnError(func(r *colly.Response, err error) {
		e = err
	})
	c.Visit("https://" + url)

	return info, e
}
