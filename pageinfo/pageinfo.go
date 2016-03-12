package pageinfo

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const maxResponseBodySize = 1024 * 1024 * 5

type PageInfo struct {
	ResponseBody []byte
	ContentType  string
	StatusCode   int
	URL          *url.URL
	ForeignHosts []string
}

func New(r *http.Response) (*PageInfo, error) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, maxResponseBodySize))
	return &PageInfo{
		body,
		r.Header.Get("Content-Type"),
		r.StatusCode,
		r.Request.URL,
		make([]string, 0),
	}, err
}

func (p *PageInfo) IsNewForeignHost(u *url.URL) bool {
	if u.Host == "" || u.Host == p.URL.Host {
		return false
	}
	for _, hostName := range p.ForeignHosts {
		if hostName == u.Host {
			return false
		}
	}
	p.ForeignHosts = append(p.ForeignHosts, u.Host)
	return true
}
