package checker

import (
	"github.com/asciimoo/privacyscore/penalty"
	"github.com/asciimoo/privacyscore/result"
)

type HTTPSChecker struct{}

func (c *HTTPSChecker) Check(r *result.Result) {
	if r.URL.Scheme != "https" {
		r.AddPenalty(penalty.P_NO_HTTPS, 5)
	}
}
