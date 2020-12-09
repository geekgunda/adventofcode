package main

import "testing"

func TestDay01FindMatchingPair(t *testing.T) {
	var input = []int64{1721,
		979,
		366,
		299,
		675,
		1456}
	targetSum := int64(2020)
	targetProduct := int64(514579)
	res, err := FindMatchingPair(input, targetSum)
	if err != nil {
		t.Errorf("Error : %v", err)
	}
	if res != targetProduct {
		t.Errorf("Invalid result: %v", res)
	}
}

func TestDay01FindMatchingTriplet(t *testing.T) {
	var input = []int64{1721,
		979,
		366,
		299,
		675,
		1456}
	targetSum := int64(2020)
	targetProduct := int64(241861950)
	res, err := FindMatchingTriplet(input, targetSum)
	if err != nil {
		t.Errorf("Error : %v", err)
	}
	if res != targetProduct {
		t.Errorf("Invalid result: %v", res)
	}
}
