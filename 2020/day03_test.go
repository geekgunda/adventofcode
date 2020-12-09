package main

import "testing"

func TestDay03GetTreeCount(t *testing.T) {
	input := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
	count := GetTreeCount(input, 1, 3)
	if count != 7 {
		t.Errorf("Invalid count: %d", count)
	}
}

func TestDay03GetMultipleTreeCount(t *testing.T) {
	input := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}
	slopes := [][]int{
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{2, 1},
	}
	product := GetMultipleTreeCount(input, slopes)
	if product != int64(336) {
		t.Errorf("Invalid product: %d", product)
	}
}
