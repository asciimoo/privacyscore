package checker

import (
	"github.com/asciimoo/privacyscore/penalty"
)

type CookieChecker struct{}

func (_ *CookieChecker) Check(c *CheckJob, p *PageInfo) {
	if len(p.Cookies) > 0 {
		c.Result.Penalties.Add(penalty.P_COOKIE)
	}
}
