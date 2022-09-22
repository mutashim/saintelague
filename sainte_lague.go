package saintelague

func Calculate(votes []Vote, seats int, threshold float64) *[]Vote {

	// STEP 1: Calculate Total Valid Votes and clean other data
	total_votes := 0
	for i, v := range votes {
		total_votes += v.VotesGained
		votes[i].SeatGained = 0
		votes[i].Passed = false
		votes[i].Quotient = 0
	}

	// STEP 2: Check Parliamentary Threshold
	for i, v := range votes {
		if float64(v.VotesGained) < threshold*float64(total_votes) {
			votes[i].Passed = false
			votes[i].Quotient = 0
		} else {
			votes[i].Passed = true
			votes[i].Quotient = float64(v.VotesGained)
		}
	}

	// STEP 3: Allocate
	seats_allocated := 0
	for seats_allocated < seats {

		max_index := 0
		max_value := 0.0
		for i, v := range votes {
			if v.Quotient > float64(max_value) {
				max_index = i
				max_value = v.Quotient
			}
		}

		// Allocate one seat
		votes[max_index].SeatGained++
		votes[max_index].Quotient = float64(Quotient(votes[max_index].VotesGained, votes[max_index].SeatGained))
		seats_allocated++
	}

	return &votes
}
