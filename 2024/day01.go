package main

import (
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := "1.txt"
	bytes, err := os.ReadFile(input)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	var left, right []int
	var size int
	for line := range strings.Lines(string(bytes)) {
		size++
		locIDs := strings.Split(line, "   ")
		if len(locIDs) != 2 {
			log.Fatalf("Parsed line doesn't have 2 elements: %#v", locIDs)
		}
		if loc, err := strconv.Atoi(locIDs[0]); err != nil {
			log.Fatalf("Failed parsing index: %v | location ID: %v | Err: %v", size, locIDs[0], err)
		} else {
			left = append(left, loc)
		}
		if loc, err := strconv.Atoi(strings.Trim(locIDs[1], "\n")); err != nil {
			log.Fatalf("Failed parsing index: %v | location ID: %v | Err: %v", size, locIDs[1], err)
		} else {
			right = append(right, loc)
		}
	}
	slices.Sort(left)
	slices.Sort(right)
	var result int
	for i := 0; i < size; i++ {
		result += int(math.Abs(float64(left[i] - right[i])))
	}
	log.Printf("Result for part 1: %v", result)
}
