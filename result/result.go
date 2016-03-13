package result

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"

	"github.com/asciimoo/privacyscore/penalty"

	"golang.org/x/net/publicsuffix"
)

const maxResponseBodySize = 1024 * 1024 * 5

type Result struct {
	Penalties    []*penalty.Penalty
	Errors       []error
	Score        penalty.Score
	ResponseBody []byte
	ContentType  string
	StatusCode   int
	URL          *url.URL
	ForeignHosts []string
	OriginalURL  *url.URL
	Cookies      []*http.Cookie
	Domain       string
}

var baseScore penalty.Score = 100
var mutex = &sync.Mutex{}

func New(URL string, r *http.Response) (*Result, error) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, maxResponseBodySize))
	u, _ := url.Parse(URL)
	return &Result{
		make([]*penalty.Penalty, 0),
		make([]error, 0),
		baseScore,
		body,
		r.Header.Get("Content-Type"),
		r.StatusCode,
		r.Request.URL,
		make([]string, 0),
		u,
		r.Cookies(),
		CropSubdomains(u.Host),
	}, err
}

func (r *Result) AddError(e error) {
	mutex.Lock()
	r.Errors = append(r.Errors, e)
	mutex.Unlock()
}

func (r *Result) AddPenalty(desc string, s penalty.Score) {
	p := penalty.New(desc, s)
	mutex.Lock()
	r.Score -= p.Value
	r.Penalties = append(r.Penalties, p)
	mutex.Unlock()
}

func (r *Result) IsNewForeignHost(u *url.URL) bool {
	if u.Host == "" {
		return false
	}
	host := CropSubdomains(u.Host)
	if host == r.Domain {
		return false
	}
	for _, hostName := range r.ForeignHosts {
		if hostName == host {
			return false
		}
	}
	mutex.Lock()
	r.ForeignHosts = append(r.ForeignHosts, host)
	mutex.Unlock()
	return true
}

func CropSubdomains(domain string) string {
	host, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		return domain
	}
	return host
}

func (r *Result) GetScoreName() string {
	switch {
	case r.Score >= 80:
		return "good"
	case r.Score >= 50:
		return "medium"
	case r.Score >= 0:
		return "bad"
	}
	return "poor"
}
