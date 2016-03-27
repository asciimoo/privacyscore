package checker

import (
	"fmt"
	"strings"

	"github.com/asciimoo/css/scanner"
)

type CSSChecker struct{}

func (_ *CSSChecker) Check(c *CheckJob, p *PageInfo) {
	if !strings.Contains(strings.ToLower(p.ContentType), "css") {
		return
	}
	s := scanner.New(string(p.ResponseBody))
	for {
		t := s.Next()
		switch t.Type {
		case scanner.TokenEOF:
			return
		case scanner.TokenURI:
			if strings.HasPrefix(t.Value, "data:") {
				break
			}
			var url string
			if strings.HasPrefix(t.Value, "url(") {
				if len(t.Value) < 4 {
					break
				}
				switch t.Value[4] {
				case byte('"'), byte('\''):
					url = t.Value[5 : len(t.Value)-2]
				default:
					url = t.Value[4 : len(t.Value)-1]
				}
			} else {
				url = t.Value
			}
			fmt.Println(url)
		case scanner.TokenIncludes:
			// TODO
			break
		}
	}
}
