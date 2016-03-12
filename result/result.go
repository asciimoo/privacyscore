package result

import (
	"github.com/asciimoo/privacyscore/penalty"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
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
	if u.Host == "" || u.Host == r.URL.Host {
		return false
	}
	for _, hostName := range r.ForeignHosts {
		if hostName == u.Host {
			return false
		}
	}
	mutex.Lock()
	r.ForeignHosts = append(r.ForeignHosts, u.Host)
	mutex.Unlock()
	return true
}

func (r *Result) GetScoreName() string {
	switch {
	case r.Score > 80:
		return "good"
	case r.Score > 60:
		return "medium"
	case r.Score > 0:
		return "bad"
	}
	return "poor"
}
