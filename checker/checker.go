package checker

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/asciimoo/privacyscore/result"
	"github.com/asciimoo/privacyscore/utils"
)

const USER_AGENT = "Mozilla/5.0 (compatible) PrivacyScore Checker v0.1.0"
const TIMEOUT = 5
const maxResponseBodySize = 1024 * 1024 * 5

var mutex = &sync.Mutex{}

type Checker interface {
	Check(*result.Result, *PageInfo)
}

type Check struct {
	sync.RWMutex
	Result    *result.Result
	Resources map[string]*PageInfo
}

type PageInfo struct {
	ResponseBody   []byte
	ContentType    string
	StatusCode     int
	URL            *url.URL
	OriginalURL    *url.URL
	Cookies        []*http.Cookie
	Domain         string
	ResponseHeader *http.Header
}

var checkers []Checker = []Checker{
	&CookieChecker{},
	&HTMLChecker{},
	&HTTPSChecker{},
	&SecureHeaderChecker{},
}

func Run(URL string) (*result.Result, error) {
	if !strings.HasPrefix(URL, "http://") && !strings.HasPrefix(URL, "https://") {
		URL = "http://" + URL
	}
	c := newCheck(URL)
	_, err := c.CheckURL(URL)
	return c.Result, err
}

func newCheck(URL string) *Check {
	return &Check{
		Result:    result.New(URL),
		Resources: make(map[string]*PageInfo),
	}
}

func (c *Check) CheckURL(URL string) (*PageInfo, error) {
	if u, found := c.Resources[URL]; found {
		return u, errors.New("URL already added")
	}
	r, err := fetchURL(URL)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, maxResponseBodySize))
	u, _ := url.Parse(URL)
	p := &PageInfo{
		body,
		r.Header.Get("Content-Type"),
		r.StatusCode,
		r.Request.URL,
		u,
		r.Cookies(),
		utils.CropSubdomains(r.Request.URL.Host),
		&r.Header,
	}
	if err != nil {
		c.Result.AddError(err)
	}
	c.Lock()
	c.Resources[URL] = p
	c.Unlock()
	for _, ch := range checkers {
		ch.Check(c.Result, p)
	}
	return p, nil
}

func fetchURL(URL string) (*http.Response, error) {
	client := http.Client{
		Timeout: time.Duration(TIMEOUT * time.Second),
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
	}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", USER_AGENT)
	response, err := client.Do(req)
	return response, err
}
