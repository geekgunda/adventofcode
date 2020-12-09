package main

import "testing"

func TestGetCustomsAnswerCount(t *testing.T) {
	var input = []string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
	}
	anyCount, everyCount := GetCustomsAnswerCount(input)
	if anyCount != 11 {
		t.Errorf("Invalid anyone answered count: %d", anyCount)
	}
	if everyCount != 6 {
		t.Errorf("Invalid everyone answered count: %d", everyCount)
	}
}
