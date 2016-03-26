package checker

import (
	"github.com/asciimoo/privacyscore/penalty"
	"github.com/asciimoo/privacyscore/result"
)

type CookieChecker struct{}

func (c *CookieChecker) Check(r *result.Result, p *PageInfo) {
	if len(p.Cookies) > 0 {
		r.Penalties.Add(penalty.P_COOKIE)
	}
}
