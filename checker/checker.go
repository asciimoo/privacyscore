package checker

import (
	"net/http"

	"github.com/asciimoo/privacyscore/result"
)

type Checker interface {
	Check(*result.Result)
}

var checkers []Checker = []Checker{
	&HTMLChecker{},
}

func Run(URL string) (*result.Result, bool) {
	var r *result.Result
	response, err := http.Get(URL)
	if err != nil {
		return r, false
	}
	r, err = result.New(URL, response)
	if err != nil {
		r.AddError(err)
	}
	for _, c := range checkers {
		c.Check(r)
	}
	return r, true
}
