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
}

func New(URL string, r *http.Response) (*PageInfo, error) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, maxResponseBodySize))
	u, _ := url.Parse(URL)
	return &PageInfo{
		body,
		r.Header.Get("Content-Type"),
		r.StatusCode,
		u,
	}, err
}

func (p *PageInfo) IsSameOrigin(URL string) bool {
	u, err := url.Parse(URL)
	if err != nil {
		return false
	}
	if u.Host == "" || u.Host == p.URL.Host {
		return true
	}
	return false
}
