package main

import "testing"

func TestFindInvalidNumber(t *testing.T) {
	input := []int64{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}
	res := FindInvalidNumber(input, 5)
	if res != 127 {
		t.Errorf("Invalid result: %d", res)
	}
}
