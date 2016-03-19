package result

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"

	"github.com/asciimoo/privacyscore/penalty"
	"github.com/asciimoo/privacyscore/utils"
)

const maxResponseBodySize = 1024 * 1024 * 5

type Result struct {
	Penalties      []*penalty.Penalty
	Errors         []error
	Score          penalty.Score
	ResponseBody   []byte
	ContentType    string
	StatusCode     int
	URL            *url.URL
	OriginalURL    *url.URL
	Cookies        []*http.Cookie
	Domain         string
	ResponseHeader *http.Header
}

var baseScore penalty.Score = 100
var mutex = &sync.Mutex{}

func New(URL string, r *http.Response) *Result {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, maxResponseBodySize))
	u, _ := url.Parse(URL)
	result := &Result{
		make([]*penalty.Penalty, 0),
		make([]error, 0, 8),
		baseScore,
		body,
		r.Header.Get("Content-Type"),
		r.StatusCode,
		r.Request.URL,
		u,
		r.Cookies(),
		utils.CropSubdomains(u.Host),
		&r.Header,
	}
	if err != nil {
		result.AddError(err)
	}
	return result
}

func (r *Result) AddError(e error) {
	mutex.Lock()
	r.Errors = append(r.Errors, e)
	mutex.Unlock()
}

func (r *Result) AddPenalty(pt penalty.PenaltyType, s penalty.Score) *penalty.Penalty {
	p := penalty.New(pt, s)
	mutex.Lock()
	r.Score -= p.Value
	r.Penalties = append(r.Penalties, p)
	mutex.Unlock()
	return p
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
