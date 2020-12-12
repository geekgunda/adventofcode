package main

import (
	"sort"
	"testing"
)

func TestDay10FindJoltDifferences(t *testing.T) {
	cases := []struct {
		name        string
		input       sort.IntSlice
		joltDiff    int
		permutation int64
	}{
		{
			"c1",
			sort.IntSlice{
				16,
				10,
				15,
				5,
				1,
				11,
				7,
				19,
				6,
				12,
				4,
			},
			7 * 5,
			int64(8),
		},
		{
			"c2",
			sort.IntSlice{
				28,
				33,
				18,
				42,
				31,
				14,
				46,
				20,
				48,
				47,
				24,
				23,
				49,
				45,
				19,
				38,
				39,
				11,
				1,
				32,
				25,
				35,
				8,
				17,
				7,
				9,
				4,
				2,
				34,
				10,
				3,
			},
			22 * 10,
			int64(19208),
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			res := FindJoltDifference(tt.input)
			if res != tt.joltDiff {
				t.Errorf("Invalid jolt difference: %d", res)
			}
			perm := FindPermutations(tt.input)
			if perm != tt.permutation {
				t.Errorf("Invalid permutations: %d", perm)
			}
		})
	}
}
