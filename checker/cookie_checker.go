package checker

import (
	"github.com/asciimoo/privacyscore/result"
)

type CookieChecker struct{}

func (c *CookieChecker) Check(r *result.Result) {
	if len(r.Cookies) > 0 {
		r.AddPenalty("Automatically sets cookies", 5)
	}
}
