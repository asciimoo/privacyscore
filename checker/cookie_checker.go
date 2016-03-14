package checker

import (
	"github.com/asciimoo/privacyscore/penalty"
	"github.com/asciimoo/privacyscore/result"
)

type CookieChecker struct{}

func (c *CookieChecker) Check(r *result.Result) {
	if len(r.Cookies) > 0 {
		r.AddPenalty(penalty.P_COOKIE, 5)
	}
}
