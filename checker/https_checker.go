package checker

import (
	"github.com/asciimoo/privacyscore/result"
)

type HTTPSChecker struct{}

func (c *HTTPSChecker) Check(r *result.Result) {
	if r.URL.Scheme != "https" {
		r.AddPenalty("Uses unencrypted transport layer (no HTTPS)", 15)
	}
}
