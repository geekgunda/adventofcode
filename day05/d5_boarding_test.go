package main

import "testing"

func TestGetSeatIDFromBoardingPass(t *testing.T) {
	tt := []struct {
		name         string
		boardingPass string
		seatID       int
	}{
		{"c1", "FBFBBFFRLR", 357},
		{"c2", "BFFFBBFRRR", 567},
		{"c3", "FFFBBBFRRR", 119},
		{"c4", "BBFFBBFRLL", 820},
	}
	for _, c := range tt {
		t.Run(c.name, func(t *testing.T) {
			res := GetSeatIDFromBoardingPass(c.boardingPass)
			if res != c.seatID {
				t.Errorf("Invalid result: %d | expected: %d | pass: %s", res, c.seatID, c.boardingPass)
			}
		})
	}
}

func TestGetHighestSeatID(t *testing.T) {
	input := []string{
		"FBFBBFFRLR",
		"BFFFBBFRRR",
		"FFFBBBFRRR",
		"BBFFBBFRLL",
	}
	res := GetHighestSeatID(input)
	if res != 820 {
		t.Errorf("Invalid result: %d | expected: 820", res)
	}
}
