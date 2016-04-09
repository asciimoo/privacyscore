package scoredb

import (
	"fmt"
	"sync"

	"github.com/asciimoo/privacyscore/penalty"
)

type ScoreCount struct {
	Label     string
	Count     uint
	BaseScore penalty.Score
}

const (
	step     = 10
	minScore = 0
	maxScore = 100
)

var scores = struct {
	sync.RWMutex
	DB         []*ScoreCount
	EntryCount uint
}{DB: make([]*ScoreCount, 0, maxScore-minScore), EntryCount: 0}

func init() {
	i := minScore
	scores.DB = append(scores.DB, &ScoreCount{"< 0", 0, 0})
	for i <= maxScore {
		var label string
		switch i {
		case maxScore:
			label = fmt.Sprintf("%v", i)
		default:
			label = fmt.Sprintf("%v-%v", i, i+step-1)
		}
		scores.DB = append(scores.DB, &ScoreCount{label, 0, penalty.Score(i)})
		i += step
	}
}

func Add(s penalty.Score) {
	var idx int
	if s < minScore {
		idx = 0
	} else {
		s = s - (s % step)
		idx = int((int(s)-minScore)/step) + 1
	}
	if idx >= len(scores.DB) {
		idx = len(scores.DB) - 1
	}
	scores.Lock()
	scores.EntryCount += 1
	scores.DB[idx].Count += 1
	scores.Unlock()
}

func GetAll() []*ScoreCount {
	return scores.DB
}

func GetTopEntryCount() uint {
	var c uint = 0
	for _, e := range scores.DB {
		if e.Count > c {
			c = e.Count
		}
	}
	return c
}
