package whois

import (
	"strings"
	"time"

	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

type WhoisInfo struct {
	CreationDate   time.Time
	ExpirationDate time.Time
	Owner          string
	ContactEmail   string
}

func Whois(url string) (*WhoisInfo, error) {
	raw, err := whois.Whois(strings.Replace(url, "www", "", 1))
	if err != nil {
		return nil, err
	}
	result, err := whoisparser.Parse(raw)
	if err != nil {
		return nil, err
	}
	return &WhoisInfo{
		CreationDate:   *result.Domain.CreatedDateInTime,
		ExpirationDate: *result.Domain.ExpirationDateInTime,
		Owner:          result.Registrant.Name,
		ContactEmail:   result.Administrative.Email,
	}, nil
}
