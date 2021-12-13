package main

import "testing"

func TestDay12CountCavePaths(t *testing.T) {
	var cases = []struct {
		input                []string
		expected, expectedV2 int
	}{
		{
			input: []string{
				"start-A",
				"start-b",
				"A-c",
				"A-b",
				"b-d",
				"A-end",
				"b-end",
			},
			expected:   10,
			expectedV2: 36,
		},
		{
			input: []string{
				"dc-end",
				"HN-start",
				"start-kj",
				"dc-start",
				"dc-HN",
				"LN-dc",
				"HN-end",
				"kj-sa",
				"kj-HN",
				"kj-dc",
			},
			expected:   19,
			expectedV2: 103,
		},
		{
			input: []string{
				"fs-end",
				"he-DX",
				"fs-he",
				"start-DX",
				"pj-DX",
				"end-zg",
				"zg-sl",
				"zg-pj",
				"pj-he",
				"RW-he",
				"fs-DX",
				"pj-RW",
				"zg-RW",
				"start-pj",
				"he-WI",
				"zg-he",
				"pj-fs",
				"start-RW",
			},
			expected:   226,
			expectedV2: 3509,
		},
	}
	for i, tt := range cases {
		paths, pathsV2 := countCavePaths(tt.input)
		if paths != tt.expected {
			t.Errorf("Path count mis-match: %d | case: %d", paths, i)
		}
		if pathsV2 != tt.expectedV2 {
			t.Errorf("Path count mis-match for double-traversal: %d | case: %d", pathsV2, i)
		}
	}
}
