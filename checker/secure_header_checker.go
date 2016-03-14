package checker

import (
	"github.com/asciimoo/privacyscore/penalty"
	"github.com/asciimoo/privacyscore/result"
)

type SecureHeaderChecker struct{}

func (c *SecureHeaderChecker) Check(r *result.Result) {
	missingSecureHeaders := make([]string, 0, 8)
	if r.ResponseHeader.Get("X-Xss-Protection") != "1; mode=block" {
		missingSecureHeaders = append(missingSecureHeaders, "X-Xss-Protection")
	}
	if r.ResponseHeader.Get("X-Content-Type-Options") != "nosniff" {
		missingSecureHeaders = append(missingSecureHeaders, "X-Content-Type-Options")
	}
	if len(missingSecureHeaders) > 0 {
		// TODO scoring
		p := r.AddPenalty(penalty.P_NO_SECURE_HEADER, penalty.Score(len(missingSecureHeaders)*3))
		p.Notes = missingSecureHeaders
	}
}
