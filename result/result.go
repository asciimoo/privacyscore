package result

import (
	"github.com/asciimoo/privacyscore/penalty"
)

type Result struct {
	Penalties []*penalty.Penalty
	Errors    []error
	Score     penalty.Score
}

var baseScore penalty.Score = 100

func New() *Result {
	return &Result{make([]*penalty.Penalty, 0), make([]error, 0), baseScore}
}

func (r *Result) AddError(e error) {
	r.Errors = append(r.Errors, e)
}

func (r *Result) AddPenalty(desc string, s penalty.Score) {
	p := penalty.New(desc, s)
	r.Score -= p.Value
	r.Penalties = append(r.Penalties, p)
}
