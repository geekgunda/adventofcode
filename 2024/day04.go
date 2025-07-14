//go:build ignore

package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	input := "4.txt"
	bytes, err := os.ReadFile(input)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	var words []string
	for line := range strings.Lines(string(bytes)) {
		words = append(words, strings.Trim(line, "\n"))
	}
	count := wordSearch(words, "XMAS")
	log.Printf("Part 1 | Match count: %v", count)
	xCount := doubleDiagonalSearch(words, "MAS")
	log.Printf("Part 2 | Match count: %v", xCount)
}

func doubleDiagonalSearch(matrix []string, word string) int {
	var matched int
	wl := len(word)
	// The center of the word where the X pattern matches
	wc := 1
	// Loop over the 2-D matrix
	for i := range matrix {
		for j := range len(matrix[i]) {
			// Move on if center word doesn't match
			if matrix[i][j] != word[wc] {
				continue
			}
			// Track match count for this position
			match := 0
			// Loop to check if diagonals match
			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					// We only care about diagonals
					if x == 0 || y == 0 {
						continue
					}
					wk := 0
					// Start from the beginning of the word
					wi := i + x*wc
					wj := j + y*wc
					for wk < wl {
						// Exit if out of bound
						if wi < 0 || wi >= len(matrix) || wj < 0 || wj >= len(matrix[i]) {
							break
						}
						// Exit if word stops matching
						if matrix[wi][wj] != word[wk] {
							break
						}
						// Move to comparing next character in the word
						wk++
						// We go up or down diagonally!
						wi -= x
						wj -= y
					}
					if wk == wl {
						//log.Printf("Matched index: [%v][%v]", i, j)
						match++
					}
				}
			}
			// Need both diagonals to match, for it to count!
			if match >= 2 {
				matched++
			}
		}
	}

	return matched
}

func wordSearch(matrix []string, word string) int {
	var matched int

	wl := len(word)
	// Loop to traverse the 2-D matrix
	for i := range matrix {
		for j := range len(matrix[i]) {
			// Move on if first word doesn't match
			if matrix[i][j] != word[0] {
				continue
			}
			// Loop to traverse in all 8 directions
			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					// Ignore the starting position
					if x == 0 && y == 0 {
						continue
					}
					// Iter variable to compare against the word
					wk := 1
					wi := i + x
					wj := j + y
					// Loop to search for word in a given direction
					for wk < wl {
						// Exit if out of bound
						if wi < 0 || wi >= len(matrix) || wj < 0 || wj >= len(matrix[i]) {
							break
						}
						// Exit if word stops matching
						if matrix[wi][wj] != word[wk] {
							break
						}
						// Move to comparing next character in the word
						wk++
						wi += x
						wj += y
					}
					if wk == wl {
						matched++
					}
				}
			}
		}
	}
	return matched
}
