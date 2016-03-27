package checker

import (
	"bytes"
	"errors"
	"io"
	"net/url"
	"strings"

	"golang.org/x/net/html"

	"github.com/asciimoo/privacyscore/penalty"
	"github.com/asciimoo/privacyscore/utils"
)

type HTMLChecker struct{}

func (_ *HTMLChecker) Check(c *CheckJob, p *PageInfo) {
	if !strings.Contains(strings.ToLower(p.ContentType), "html") {
		return
	}
	forbidsReferrer := false
	hasHTTPLink := false
	scriptTagFound := false
	t := html.NewTokenizer(bytes.NewReader(p.ResponseBody))
	for {
		tagToken := t.Next()
		if tagToken == html.ErrorToken {
			if t.Err() == io.EOF {
				break
			} else {
				c.Result.AddError(errors.New("Invalid HTML content"))
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
				c.Result.Penalties.Add(penalty.P_JS)
				scriptTagFound = true
			}
			src, found := getAttr(t, "src")
			if found {
				u, err := url.Parse(src)
				if err != nil {
					break
				}
				if isForeignHost(u.Host, p.Domain) {
					c.Result.Penalties.Add(penalty.P_EXTERNAL_RESOURCE, utils.CropSubdomains(u.Host))
				}
			}
		case "iframe":
			src, found := getAttr(t, "src")
			if found {
				u, err := url.Parse(src)
				if err == nil && isForeignHost(u.Host, p.Domain) {
					c.Result.Penalties.Add(penalty.P_IFRAME, utils.CropSubdomains(u.Host))
				}
			}
		case "link":
			attrs := getAttrs(t)
			if rel, found := attrs["rel"]; !found || rel != "stylesheet" {
				break
			}
			if src, found := attrs["href"]; found {
				u, err := url.Parse(src)
				if err == nil && isForeignHost(u.Host, p.Domain) {
					c.Result.Penalties.Add(penalty.P_EXTERNAL_RESOURCE, utils.CropSubdomains(u.Host))
				}
			}
		case "img":
			src, found := getAttr(t, "src")
			if !found {
				break
			}
			u, err := url.Parse(src)
			if err == nil && isForeignHost(u.Host, p.Domain) {
				c.Result.Penalties.Add(penalty.P_EXTERNAL_RESOURCE, utils.CropSubdomains(u.Host))
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
			if !forbidsReferrer && !noreferrer && isForeignHost(u.Host, p.Domain) {
				c.Result.Penalties.Add(penalty.P_EXTERNAL_LINK, utils.CropSubdomains(u.Host))
			}
		}
	}
	if hasHTTPLink {
		c.Result.Penalties.Add(penalty.P_HTTP_LINK)
	}
}

func isForeignHost(host, baseDomain string) bool {
	host = utils.CropSubdomains(host)
	if host == "" || host == baseDomain {
		return false
	}
	return true
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
