package checker

import (
	"bytes"
	"errors"
	"io"
	"net/url"
	"strings"

	"golang.org/x/net/html"

	"github.com/asciimoo/privacyscore/penalty"
	"github.com/asciimoo/privacyscore/result"
	"github.com/asciimoo/privacyscore/utils"
)

type HTMLChecker struct{}

func (c *HTMLChecker) Check(r *result.Result) {
	if !strings.Contains(strings.ToLower(r.ContentType), "html") {
		r.AddError(errors.New("No HTML content found"))
		return
	}
	scriptTagFound := false
	forbidsReferrer := false
	externalLinkHosts := make([]string, 0, 8)
	externalResourceHosts := make([]string, 0, 8)
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
				r.AddPenalty(penalty.P_JS, 5)
				scriptTagFound = true
			}
			src, found := getAttr(t, "src")
			if found {
				u, err := url.Parse(src)
				if err != nil {
					break
				}
				addHostIfNew(u.Host, r.Domain, &externalResourceHosts)
			}
		case "link":
			attrs := getAttrs(t)
			if rel, found := attrs["rel"]; !found || rel != "stylesheet" {
				break
			}
			if src, found := attrs["href"]; found {
				u, err := url.Parse(src)
				if err != nil {
					break
				}
				addHostIfNew(u.Host, r.Domain, &externalResourceHosts)
			}
		case "img":
			src, found := getAttr(t, "src")
			if !found {
				break
			}
			u, err := url.Parse(src)
			if err != nil {
				break
			}
			addHostIfNew(u.Host, r.Domain, &externalResourceHosts)
		case "meta":
			attrs := getAttrs(t)
			if _, found := attrs["name"]; !found || attrs["name"] != "referrer" {
				break
			}
			if _, found := attrs["content"]; !found {
				break
			}
			switch strings.ToLower(attrs["content"]) {
			case "never", "none", "origin", "no-referrer":
				forbidsReferrer = true
			}
		case "a":
			attrs := getAttrs(t)
			src, found := attrs["href"]
			if !found {
				break
			}
			noreferrer := false
			if rel, found := attrs["rel"]; found && rel == "noreferrer" {
				noreferrer = true
			}
			u, err := url.Parse(src)
			if err != nil {
				break
			}
			if (u.Scheme == "" && r.URL.Scheme != "https") || u.Scheme == "http" {
				hasHTTPLink = true
			}
			if !forbidsReferrer && !noreferrer {
				addHostIfNew(u.Host, r.Domain, &externalLinkHosts)
			}
		}
	}
	if len(externalLinkHosts) > 0 {
		p := r.AddPenalty(penalty.P_EXTERNAL_LINK, 2)
		p.Notes = externalLinkHosts
	}
	if hasHTTPLink {
		r.AddPenalty(penalty.P_HTTP_LINK, 2)
	}
	if len(externalResourceHosts) > 0 {
		p := r.AddPenalty(penalty.P_EXTERNAL_RESOURCE, penalty.Score(len(externalResourceHosts)*10))
		p.Notes = externalResourceHosts
	}
}

func addHostIfNew(host, self string, a *[]string) {
	host = utils.CropSubdomains(host)
	if host == "" || host == self {
		return
	}
	for _, h := range *a {
		if h == host {
			return
		}
	}
	*a = append(*a, host)
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
