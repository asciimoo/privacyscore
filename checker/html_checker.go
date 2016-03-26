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

func (c *HTMLChecker) Check(r *result.Result, p *PageInfo) {
	if !strings.Contains(strings.ToLower(p.ContentType), "html") {
		r.AddError(errors.New("No HTML content found"))
		return
	}
	forbidsReferrer := false
	hasHTTPLink := false
	scriptTagFound := false
	externalIFrameHosts := make([]string, 0, 8)
	externalLinkHosts := make([]string, 0, 8)
	externalResourceHosts := make([]string, 0, 8)
	t := html.NewTokenizer(bytes.NewReader(p.ResponseBody))
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
				r.Penalties.Add(penalty.P_JS)
				scriptTagFound = true
			}
			src, found := getAttr(t, "src")
			if found {
				u, err := url.Parse(src)
				if err != nil {
					break
				}
				addHostIfNew(u.Host, p.Domain, &externalResourceHosts)
			}
		case "iframe":
			src, found := getAttr(t, "src")
			if found {
				u, err := url.Parse(src)
				if err == nil {
					addHostIfNew(u.Host, p.Domain, &externalIFrameHosts)
				}
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
				addHostIfNew(u.Host, p.Domain, &externalResourceHosts)
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
			addHostIfNew(u.Host, p.Domain, &externalResourceHosts)
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
			if (u.Scheme == "" && p.URL.Scheme != "https") || u.Scheme == "http" {
				hasHTTPLink = true
			}
			if !forbidsReferrer && !noreferrer {
				addHostIfNew(u.Host, p.Domain, &externalLinkHosts)
			}
		}
	}
	if len(externalIFrameHosts) > 0 {
		r.Penalties.Add(penalty.P_IFRAME, externalIFrameHosts...)
	}
	if len(externalLinkHosts) > 0 {
		r.Penalties.Add(penalty.P_EXTERNAL_LINK, externalLinkHosts...)
	}
	if hasHTTPLink {
		r.Penalties.Add(penalty.P_HTTP_LINK)
	}
	if len(externalResourceHosts) > 0 {
		r.Penalties.Add(penalty.P_EXTERNAL_RESOURCE, externalResourceHosts...)
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
