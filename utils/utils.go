package utils

import (
	"golang.org/x/net/publicsuffix"
)

func CropSubdomains(domain string) string {
	host, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		return domain
	}
	return host
}
