package utils

import (
	"net/url"

	"golang.org/x/net/publicsuffix"

	"github.com/asciimoo/privacyscore/penalty"
)

func CropSubdomains(domain string) string {
	host, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		return domain
	}
	return host
}

func GetScoreName(s penalty.Score) string {
	switch {
	case s >= 80:
		return "good"
	case s >= 60:
		return "medium"
	case s >= 40:
		return "bad"
	}
	return "poor"
}

func GetFullURL(URL, baseURL *url.URL) string {
	if URL.Host == "" {
		URL.Host = baseURL.Host
	}
	switch URL.Scheme {
	case "data":
		return ""
	case "":
		URL.Scheme = baseURL.Scheme
	}
	return URL.String()
}

func IsForeignHost(host, baseDomain string) bool {
	host = CropSubdomains(host)
	if host == "" || host == baseDomain {
		return false
	}
	return true
}
