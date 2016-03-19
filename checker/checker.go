package checker

import (
	"net/http"
	"strings"
	"time"

	"github.com/asciimoo/privacyscore/result"
)

const USER_AGENT = "Mozilla/5.0 (compatible) PrivacyScore Checker v0.1.0"
const TIMEOUT = 5

type Checker interface {
	Check(*result.Result)
}

var checkers []Checker = []Checker{
	&CookieChecker{},
	&HTMLChecker{},
	&HTTPSChecker{},
	&SecureHeaderChecker{},
}

func Run(URL string) (*result.Result, error) {
	if !strings.HasPrefix(URL, "http://") && !strings.HasPrefix(URL, "https://") {
		URL = "http://" + URL
	}
	var r *result.Result
	client := http.Client{
		Timeout: time.Duration(TIMEOUT * time.Second),
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
	}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return r, err
	}
	req.Header.Set("User-Agent", USER_AGENT)
	response, err := client.Do(req)
	if err != nil {
		return r, err
	}
	defer response.Body.Close()
	r = result.New(URL, response)
	for _, c := range checkers {
		c.Check(r)
	}
	return r, nil
}
