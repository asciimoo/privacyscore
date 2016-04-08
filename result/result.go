package result

import (
	"sync"

	"github.com/asciimoo/privacyscore/penalty"
)

type Result struct {
	sync.RWMutex
	Penalties *penalty.PenaltyContainer
	Errors    []error
	BaseURL   string
}

func New(URL string) *Result {
	return &Result{
		Penalties: penalty.NewPenaltyContainer(),
		Errors:    make([]error, 0, 8),
		BaseURL:   URL,
	}
}

func (r *Result) AddError(e error) {
	for _, e2 := range r.Errors {
		if e.Error() == e2.Error() {
			return
		}
	}
	r.Lock()
	r.Errors = append(r.Errors, e)
	r.Unlock()
}
