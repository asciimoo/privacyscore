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
	forbidsReferrer := false
	hasExternalLink := false
	hasHTTPLink := false
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
		case "img":
			src, found := getAttr(t, "src")
			if !found {
				break
			}
			u, _ := url.Parse(src)
			if r.IsNewForeignHost(u) {
				r.AddPenalty("Loads external resource from "+u.Host, 10)
			}
		case "meta":
			attrs := getAttrs(t)
			if _, found := attrs["name"]; !found || attrs["name"] != "referrer" {
				break
			}
			if _, found := attrs["content"]; !found {
				break
			}
			switch strings.ToLower(attrs["content"]) {
			case "never", "none", "origin":
				forbidsReferrer = true
			}
		case "a":
			src, found := getAttr(t, "href")
			if !found {
				break
			}
			u, _ := url.Parse(src)
			if (u.Scheme == "" && r.URL.Scheme != "https") || u.Scheme == "http" {
				hasHTTPLink = true
			}
			if forbidsReferrer || hasExternalLink {
				break
			}
			if r.IsNewForeignHost(u) {
				hasExternalLink = true
			}
		}
	}
	if hasExternalLink && !forbidsReferrer {
		r.AddPenalty("Has link to foreign host without HTTP referrer restrictions", 10)
	}
	if hasHTTPLink {
		r.AddPenalty("Has link to unencrypted host", 2)
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
