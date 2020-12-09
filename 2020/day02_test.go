package main

import "testing"

func TestDay02GetValidPasswordCountByPart1Policy(t *testing.T) {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
	ct, err := GetValidPasswordCount(input, false)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if ct != 2 {
		t.Errorf("Invalid count: %d", ct)
	}
}

func TestDay02GetValidPasswordCountByPart2Policy(t *testing.T) {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
	ct, err := GetValidPasswordCount(input, true)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if ct != 1 {
		t.Errorf("Invalid count: %d", ct)
	}
}
