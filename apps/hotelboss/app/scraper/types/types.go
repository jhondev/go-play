package types

import "time"

type HotelSiteInfo struct {
	LogoURL string
}
type Endpoint struct {
	Name string
}
type HostInfo struct {
	Protocol       string
	CreationDate   time.Time
	ExpirationDate time.Time
	Owner          string
	ContactEmail   string
	Endpoints      []Endpoint
}
type HotelInfo struct {
	Site *HotelSiteInfo
	Host *HostInfo
}

type ScrapeHotelFunc func(url string) (*HotelSiteInfo, error)
