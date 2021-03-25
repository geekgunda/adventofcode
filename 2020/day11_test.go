package main

import "testing"

func TestDay11FindStableLayout(t *testing.T) {
	input := []string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}
	res := FindStableLayoutSeatCount(input, 4, 1)
	if res != 37 {
		t.Errorf("Invalid seat count: %d\n", res)
	}
	res2 := FindStableLayoutSeatCount(input, 5, 0)
	if res2 != 26 {
		t.Errorf("Invalid seat count (Part 2): %d\n", res2)
	}
}
