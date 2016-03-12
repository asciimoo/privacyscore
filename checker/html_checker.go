package checker

import (
	"bytes"
	"errors"
	"io"
	"net/url"
	"strings"

	"golang.org/x/net/html"

	"github.com/asciimoo/privacyscore/result"
)

type HTMLChecker struct{}

func (c *HTMLChecker) Check(r *result.Result) {
	if !strings.Contains(strings.ToLower(r.ContentType), "html") {
		r.AddError(errors.New("No HTML content found"))
		return
	}
	scriptTagFound := false
	t := html.NewTokenizer(bytes.NewReader(r.ResponseBody))
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
		if tagToken != html.StartTagToken && tagToken != html.SelfClosingTagToken {
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
			if !found {
				break
			}
			u, _ := url.Parse(src)
			if r.IsNewForeignHost(u) {
				r.AddPenalty("Loads external resource from "+u.Host, 10)
			}
		case "link":
			attrs := getAttrs(t)
			if _, found := attrs["href"]; !found {
				break
			}
			u, _ := url.Parse(attrs["href"])
			if r.IsNewForeignHost(u) {
				r.AddPenalty("Loads external resource from "+u.Host, 10)
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

func getAttrs(t *html.Tokenizer) map[string]string {
	attrs := make(map[string]string)
	for {
		attrName, attrValue, more := t.TagAttr()
		attrs[string(attrName)] = string(attrValue)
		if !more {
			break
		}
	}
	return attrs
}
