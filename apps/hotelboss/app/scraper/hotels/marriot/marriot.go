package marriot

import (
	"hotelboss/app/scraper/types"

	"github.com/gocolly/colly/v2"
)

const URL = "www.marriott.com"

func Scrape(url string) (*types.HotelSiteInfo, error) {
	var e error
	info := &types.HotelSiteInfo{}

	c := colly.NewCollector()
	// Find and get icon link
	c.OnHTML("link[rel=\"shortcut icon\"]", func(e *colly.HTMLElement) {
		info.LogoURL = e.Attr("href")
	})
	c.OnError(func(r *colly.Response, err error) {
		e = err
	})
	c.Visit("https://" + url)

	return info, e
}
