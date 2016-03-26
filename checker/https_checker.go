package checker

import (
	"github.com/asciimoo/privacyscore/penalty"
	"github.com/asciimoo/privacyscore/result"
)

type HTTPSChecker struct{}

func (c *HTTPSChecker) Check(r *result.Result, p *PageInfo) {
	if p.URL.Scheme != "https" {
		r.Penalties.Add(penalty.P_NO_HTTPS)
	}
}
