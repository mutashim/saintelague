package saintelague

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	votes := []Vote{
		{PartyName: "A", VotesGained: 30000},
		{PartyName: "B", VotesGained: 5000},
		{PartyName: "C", VotesGained: 15000},
		{PartyName: "D", VotesGained: 20000},
		{PartyName: "E", VotesGained: 7500},
		{PartyName: "F", VotesGained: 2500},
	}

	expected := []Vote{
		{PartyName: "A", VotesGained: 30000, SeatGained: 3, Quotient: 4285.714285714285, Passed: true},
		{PartyName: "B", VotesGained: 5000, SeatGained: 0, Quotient: 5000, Passed: true},
		{PartyName: "C", VotesGained: 15000, SeatGained: 1, Quotient: 5000, Passed: true},
		{PartyName: "D", VotesGained: 20000, SeatGained: 2, Quotient: 4000, Passed: true},
		{PartyName: "E", VotesGained: 7500, SeatGained: 1, Quotient: 2500, Passed: true},
		{PartyName: "F", VotesGained: 2500, SeatGained: 0, Quotient: 0, Passed: false},
	}

	result := Calculate(votes, 7, 0.04)

	assert.Equal(t, expected, *result)
}
