package checker

import (
	"bytes"
	"errors"
	"io"
	"strings"

	"golang.org/x/net/html"

	"github.com/asciimoo/privacyscore/pageinfo"
	"github.com/asciimoo/privacyscore/result"
)

type HTMLChecker struct{}

func (c *HTMLChecker) Check(info *pageinfo.PageInfo, r *result.Result) {
	if !strings.Contains(strings.ToLower(info.ContentType), "html") {
		r.AddError(errors.New("No HTML content found"))
		return
	}
	scriptTagFound := false
	t := html.NewTokenizer(bytes.NewReader(info.ResponseBody))
	for {
		tagToken := t.Next()
		if tagToken == html.ErrorToken {
			if t.Err() == io.EOF {
				break
			} else {
				r.AddError(errors.New("Invalid HTML content"))
				return
			}
		}
		if tagToken != html.StartTagToken {
			continue
		}
		tagName, _ := t.TagName()
		switch string(tagName) {
		case "script":
			if !scriptTagFound {
				r.AddPenalty("Uses javascript", 5)
				scriptTagFound = true
			}
			src, found := getAttr(t, "src")
			if found && !info.IsSameOrigin(src) {
				r.AddPenalty("Loads external javascript resource: "+string(src), 10)
			}
		}
	}
}

func getAttr(t *html.Tokenizer, name string) (string, bool) {
	for {
		attrName, attrValue, more := t.TagAttr()
		if string(attrName) == name {
			return string(attrValue), true
		}
		if !more {
			break
		}
	}
	return "", false
}
