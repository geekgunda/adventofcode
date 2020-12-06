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
	count := GetCustomsAnswerCount(input)
	if count != 11 {
		t.Errorf("Invalid count: %d", count)
	}
}
