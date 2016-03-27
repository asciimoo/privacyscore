package checker

import (
	"github.com/asciimoo/privacyscore/penalty"
)

type SecureHeaderChecker struct{}

func (_ *SecureHeaderChecker) Check(c *CheckJob, p *PageInfo) {
	switch p.ResponseHeader.Get("X-Frame-Options") {
	case "DENY", "SAMEORIGIN":
		break
	default:
		c.Result.Penalties.Add(penalty.P_NO_SECURE_HEADER, "X-Frame-Options")
	}
	if p.ResponseHeader.Get("X-Xss-Protection") != "1; mode=block" {
		c.Result.Penalties.Add(penalty.P_NO_SECURE_HEADER, "X-Xss-Protection")
	}
	if p.ResponseHeader.Get("X-Content-Type-Options") != "nosniff" {
		c.Result.Penalties.Add(penalty.P_NO_SECURE_HEADER, "X-Content-Type-Options")
	}
	if p.URL.Scheme == "https" && p.ResponseHeader.Get("Strict-Transport-Security") == "" {
		c.Result.Penalties.Add(penalty.P_NO_SECURE_HEADER, "Strict-Transport-Security")
	}
}
