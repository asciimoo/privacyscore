package checker

import (
	"net/http"

	"github.com/asciimoo/privacyscore/pageinfo"
	"github.com/asciimoo/privacyscore/result"
)

type Checker interface {
	Check(*pageinfo.PageInfo, *result.Result)
}

var checkers []Checker = []Checker{
	&HTMLChecker{},
}

func Run(URL string) (checkResult *result.Result) {
	checkResult = result.New()
	response, err := http.Get(URL)
	if err != nil {
		checkResult.AddError(err)
		return
	}
	info, err := pageinfo.New(response)
	if err != nil {
		checkResult.AddError(err)
	}
	for _, c := range checkers {
		c.Check(info, checkResult)
	}
	return
}
