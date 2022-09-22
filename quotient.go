package saintelague

func Quotient(votesGained int, seatGained int) float64 {
	return (float64(votesGained) / (2*float64(seatGained) + 1))
}
