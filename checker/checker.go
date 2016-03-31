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

const RESOURCE_LIMIT = 64
const TIMEOUT = 5
const USER_AGENT = "Mozilla/5.0 (compatible) PrivacyScore Checker v0.1.0"
const maxResponseBodySize = 1024 * 1024 * 5

var mutex = &sync.Mutex{}

type Checker interface {
	Check(*CheckJob, *PageInfo)
}

type CheckJob struct {
	sync.RWMutex
	Result    *result.Result
	Resources map[string]*PageInfo
	Chan      chan bool
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
	&CSSChecker{},
}

func Run(URL string) (*CheckJob, error) {
	if !strings.HasPrefix(URL, "http://") && !strings.HasPrefix(URL, "https://") {
		URL = "http://" + URL
	}
	c := newCheckJob(URL)
	finishedResources := 0
	errorCount := 0
	c.CheckURL(URL)
	for finishedResources != len(c.Resources) && finishedResources < RESOURCE_LIMIT {
		select {
		case ret := <-c.Chan:
			if ret == false {
				errorCount += 1
			}
			finishedResources += 1
		}
	}
	if finishedResources == 0 || (errorCount > 0 && errorCount == finishedResources) {
		return c, errors.New("Could not download host")
	}
	if r, found := c.Resources[URL]; found && r != nil {
		c.Result.BaseURL = c.Resources[URL].URL.String()
	}
	return c, nil
}

func newCheckJob(URL string) *CheckJob {
	return &CheckJob{
		Result:    result.New(URL),
		Resources: make(map[string]*PageInfo),
		Chan:      make(chan bool, RESOURCE_LIMIT),
	}
}

func (c *CheckJob) CheckURL(URL string) {
	if URL == "" {
		return
	}
	// URL already added
	c.RLock()
	r, found := c.Resources[URL]
	c.RUnlock()
	if found && r != nil {
		return
	}
	// resource limit exceeded
	if len(c.Resources) >= RESOURCE_LIMIT {
		return
	}
	var p *PageInfo
	c.Lock()
	c.Resources[URL] = p
	c.Unlock()
	go func() {
		r, err := fetchURL(URL)
		if err != nil {
			c.Result.AddError(err)
			c.Chan <- false
			return
		}
		var body []byte
		contentType := r.Header.Get("Content-Type")
		if r.StatusCode == 200 && (strings.Contains(contentType, "text/") || strings.Contains(contentType, "html")) {
			body, err = ioutil.ReadAll(io.LimitReader(r.Body, maxResponseBodySize))
			if err != nil {
				c.Result.AddError(err)
			}
		} else {
			body = []byte{}
		}
		r.Body.Close()

		u, _ := url.Parse(URL)
		p = &PageInfo{
			body,
			r.Header.Get("Content-Type"),
			r.StatusCode,
			r.Request.URL,
			u,
			r.Cookies(),
			utils.CropSubdomains(r.Request.URL.Host),
			&r.Header,
		}
		c.Lock()
		c.Resources[URL] = p
		c.Unlock()
		for _, ch := range checkers {
			ch.Check(c, p)
		}
		c.Chan <- true
	}()
	return
}

func fetchURL(URL string) (*http.Response, error) {
	client := http.Client{
		Timeout: time.Duration(TIMEOUT * time.Second),
		Transport: &http.Transport{
			DisableKeepAlives:     true,
			ResponseHeaderTimeout: TIMEOUT * time.Second,
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
