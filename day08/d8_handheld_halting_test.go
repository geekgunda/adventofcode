package main

import "testing"

func TestFindAndFixLoop(t *testing.T) {
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
	insts := ParseInstructions(input)
	loopRes, _ := FindLoop(insts)
	if loopRes != 5 {
		t.Errorf("Invalid loop result: %d", loopRes)
	}
	loopRes = FixLoop(insts)
	if loopRes != 8 {
		t.Errorf("Invalid final result: %d", loopRes)
	}
}
