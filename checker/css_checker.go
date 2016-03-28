package checker

import (
	"bytes"
	"net/url"
	"regexp"
	"strings"

	"github.com/asciimoo/privacyscore/penalty"
	"github.com/asciimoo/privacyscore/utils"
)

// TODO
var URLREGEXP *regexp.Regexp = regexp.MustCompile("url\\(['\"]([\u0009\u0021\u0023-\u0026\u0028\u002a-\u007E]+)['\"]\\)")

type CSSChecker struct{}

func (_ *CSSChecker) Check(c *CheckJob, p *PageInfo) {
	if !strings.Contains(strings.ToLower(p.ContentType), "css") {
		return
	}
	for _, b := range URLREGEXP.FindAll(p.ResponseBody, 128) {
		if len(b) < 6 {
			break
		}
		switch b[4] {
		case byte('"'), byte('\''):
			b = b[5 : len(b)-2]
		default:
			b = b[4 : len(b)-1]
		}
		if bytes.HasPrefix(b, []byte("data:")) {
			break
		}
		u, err := url.Parse(string(b))
		if err != nil {
			break
		}
		if utils.IsForeignHost(u.Host, p.Domain) {
			c.Result.Penalties.Add(penalty.P_EXTERNAL_RESOURCE, utils.CropSubdomains(u.Host))
		} else {
			c.CheckURL(utils.GetFullURL(u, p.URL))
		}
	}
}
