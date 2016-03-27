package checker

import (
	"github.com/asciimoo/privacyscore/penalty"
)

type HTTPSChecker struct{}

func (_ *HTTPSChecker) Check(c *CheckJob, p *PageInfo) {
	if p.URL.Scheme != "https" {
		c.Result.Penalties.Add(penalty.P_NO_HTTPS)
	}
}
