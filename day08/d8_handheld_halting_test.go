package main

import "testing"

func TestFindLoop(t *testing.T) {
	input := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}
	res := FindLoop(input)
	if res != 5 {
		t.Errorf("Invalid result: %d", res)
	}
}
