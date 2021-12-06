package main

import (
	"fmt"
)

var cache map[string]uint64

func day6() error {
	cache = make(map[string]uint64)
	input, err := readFileAsNumbersCsv()
	if err != nil {
		return err
	}
	var count uint64
	count = simulate(input, 80)
	logResult(6, 1, "Final count after 80 days: ", count)
	fmt.Println("Cache size: ", len(cache))
	count = simulate(input, 256)
	logResult(6, 2, "Final count after 256 days: ", count)
	fmt.Println("Cache size: ", len(cache))
	return nil
}

func simulate(input []int, days int) uint64 {
	fmt.Println("Input length: ", len(input), " days: ", days)
	var count uint64
	for _, age := range input {
		val := populationCount(age+1, days)
		count += val
		//fmt.Println("item: ", i, " population: ", val)
	}
	return count
}

func populationCount(daysToReproduce, remainingDays int) uint64 {
	if daysToReproduce > remainingDays {
		return uint64(1)
	}
	// Use memoization to cache and retrieve results
	if count, ok := cache[fmt.Sprintf("%d_%d", daysToReproduce, remainingDays)]; ok {
		return count
	}
	// days taken to reproduce for current fish + it's offsprint
	count := populationCount(7, remainingDays-daysToReproduce) + populationCount(9, remainingDays-daysToReproduce)
	cache[fmt.Sprintf("%d_%d", daysToReproduce, remainingDays)] = count
	return count
}
