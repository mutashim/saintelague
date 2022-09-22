package saintelague

type Vote struct {
	PartyName   string
	VotesGained int
	SeatGained  int
	Passed      bool
	Quotient    float64
}
