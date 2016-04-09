package penalty

import (
	"sync"
)

type Score int

type PenaltyType int

const (
	P_COOKIE            PenaltyType = 1
	P_EXTERNAL_LINK     PenaltyType = 2
	P_HTTP_LINK         PenaltyType = 3
	P_EXTERNAL_RESOURCE PenaltyType = 4
	P_NO_HTTPS          PenaltyType = 5
	P_JS                PenaltyType = 6
	P_NO_SECURE_HEADER  PenaltyType = 7
	P_IFRAME            PenaltyType = 8
)

const baseScore Score = 100

type Penalty struct {
	Description string
	DetailLink  string
	Notes       []string
	value       Score
}

type PenaltyContainer struct {
	sync.RWMutex
	penalties map[PenaltyType]*Penalty
}

func NewPenaltyContainer() *PenaltyContainer {
	return &PenaltyContainer{penalties: make(map[PenaltyType]*Penalty)}
}

func (c *PenaltyContainer) GetAll() []*Penalty {
	l := make([]*Penalty, 0, len(c.penalties))
	c.RLock()
	for _, p := range c.penalties {
		l = append(l, p)
	}
	c.RUnlock()
	return l
}

func (c *PenaltyContainer) Add(pt PenaltyType, notes ...string) {
	c.RLock()
	p, found := c.penalties[pt]
	c.RUnlock()
	if found {
		for _, n := range notes {
			if n == "" {
				continue
			}
			note_found := false
			for _, pn := range p.Notes {
				if pn == n {
					note_found = true
					break
				}
			}
			if !note_found {
				c.Lock()
				p.Notes = append(p.Notes, n)
				c.Unlock()
			}
		}
	} else {
		c.Lock()
		c.penalties[pt] = New(pt)
		c.penalties[pt].Notes = notes
		c.Unlock()
	}
}

func (c *PenaltyContainer) GetScore() Score {
	score := baseScore
	for _, p := range c.penalties {
		score -= p.GetValue()
	}
	return score
}

func (p *Penalty) GetValue() Score {
	if len(p.Notes) > 0 {
		return p.value * Score(len(p.Notes))
	}
	return p.value
}

func New(p PenaltyType) *Penalty {
	desc := ""
	link := ""
	var score Score = 0
	switch p {
	case P_COOKIE:
		desc = "Automatically sets cookies"
		link = "about#p_cookie"
		score = 3
	case P_EXTERNAL_LINK:
		desc = "Sends HTTP referrer to foreign host"
		link = "about#p_external_link"
		score = 1
	case P_HTTP_LINK:
		desc = "Has link to unencrypted service (no HTTPS)"
		link = "about#p_http_link"
		score = 1
	case P_EXTERNAL_RESOURCE:
		desc = "Loads external resource"
		link = "about#p_external_resource"
		score = 5
	case P_NO_HTTPS:
		desc = "Uses unencrypted transport layer (no HTTPS)"
		link = "about#p_no_https"
		score = 6
	case P_JS:
		desc = "Uses JavaScript"
		link = "about#p_js"
		score = 7
	case P_NO_SECURE_HEADER:
		desc = "Missing secure HTTP header"
		link = "about#p_no_secure_header"
		score = 3
	case P_IFRAME:
		desc = "Loads external content to iframe"
		link = "about#p_iframe"
		score = 5
	}
	return &Penalty{desc, link, make([]string, 0, 8), score}
}
