package penalty

type Score int

type Penalty struct {
	Description string
	Value       Score
}

func New(desc string, value Score) *Penalty {
	return &Penalty{desc, value}
}
