package main

import "testing"

func TestDay14CountPolymer(t *testing.T) {
	var cases = []struct {
		input                []string
		numSteps, numStepsV2 int
		expected, expectedV2 int64
	}{
		{
			input: []string{
				"NNCB",
				"",
				"CH -> B",
				"HH -> N",
				"CB -> H",
				"NH -> C",
				"HB -> C",
				"HC -> B",
				"HN -> C",
				"NN -> C",
				"BH -> H",
				"NC -> B",
				"NB -> B",
				"BN -> B",
				"BB -> N",
				"BC -> B",
				"CC -> N",
				"CN -> C",
			},
			numSteps:   10,
			expected:   int64(1588),
			numStepsV2: 40,
			expectedV2: int64(2188189693529),
		},
	}
	for i, tt := range cases {
		//count := measurePolymerGrowth(tt.input, tt.numSteps)
		count := growPolymer(tt.input, tt.numSteps)
		if count != tt.expected {
			t.Errorf("Count mis-match: %d | case: %d", count, i)
		}
		countV2 := growPolymer(tt.input, tt.numStepsV2)
		if countV2 != tt.expectedV2 {
			t.Errorf("Count mis-match: %d | case: %d | part 2", countV2, i)
		}
	}
}
