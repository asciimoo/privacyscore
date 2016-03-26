package checker

import (
	"github.com/asciimoo/privacyscore/penalty"
	"github.com/asciimoo/privacyscore/result"
)

type SecureHeaderChecker struct{}

func (c *SecureHeaderChecker) Check(r *result.Result, p *PageInfo) {
	missingSecureHeaders := make([]string, 0, 8)
	switch p.ResponseHeader.Get("X-Frame-Options") {
	case "DENY", "SAMEORIGIN":
		break
	default:
		missingSecureHeaders = append(missingSecureHeaders, "X-Frame-Options")
	}
	if p.ResponseHeader.Get("X-Xss-Protection") != "1; mode=block" {
		missingSecureHeaders = append(missingSecureHeaders, "X-Xss-Protection")
	}
	if p.ResponseHeader.Get("X-Content-Type-Options") != "nosniff" {
		missingSecureHeaders = append(missingSecureHeaders, "X-Content-Type-Options")
	}
	if p.URL.Scheme == "https" && p.ResponseHeader.Get("Strict-Transport-Security") == "" {
		missingSecureHeaders = append(missingSecureHeaders, "Strict-Transport-Security")
	}
	if len(missingSecureHeaders) > 0 {
		r.Penalties.Add(penalty.P_NO_SECURE_HEADER, missingSecureHeaders...)
	}
}
