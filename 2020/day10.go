package main

import (
	"fmt"
	"sort"
)

func day10() error {
	input, err := readFileAsIntSlice()
	if err != nil {
		return err
	}
	res := FindJoltDifference(input)
	logResult(10, 1, "Jolt Difference (Product) is", res)
	perm := FindPermutations(input)
	logResult(10, 2, "Permutations: ", perm)
	return nil
}

func FindJoltDifference(jolts sort.IntSlice) int {
	var j1, j2, j3 int
	//fmt.Printf("Jolts origin: %v\n", jolts)
	sort.Sort(jolts)
	//fmt.Printf("Jolts sorted: %v\n", jolts)
	for i := 0; i < len(jolts); i++ {
		diff := jolts[i]
		if i > 0 {
			diff = jolts[i] - jolts[i-1]
			//fmt.Printf("n1: %d | n2: %d | diff: %d\n", jolts[i], jolts[i-1], diff)
		}
		switch diff {
		case 1:
			j1++
		case 2:
			j2++
		case 3:
			j3++
		default:
			fmt.Errorf("Invalid difference between cables: %d\n", diff)
		}
	}
	//fmt.Printf("Jolt differences are: 1j: %d | 2j: %d | 3j: %d\n", j1, j2, j3)
	j3++ // for the adapter
	return j1 * j3
}

// Count backwards to see how many ways each adapter can be connected
// Consider n+1,n+2,n+3 combinations
// Ref: https://github.com/lizthegrey/adventofcode/blob/main/2020/day10.go
func FindPermutations(jolts sort.IntSlice) int64 {
	var res int64
	lookup := make(map[int]int)
	for i, j := range jolts {
		lookup[j] = i
	}
	//fmt.Println("Input: ", jolts)
	ways := make([]int, len(jolts))
	ways[len(jolts)-1] = 1
	for i := len(jolts) - 2; i >= 0; i-- {
		sum := 0
		for diff := 1; diff <= 3; diff++ {
			if pos, ok := lookup[jolts[i]+diff]; ok {
				sum += ways[pos]
			}
		}
		//fmt.Printf("elem: %d | ways: %d\n", jolts[i], sum)
		ways[i] = sum
	}
	for v := 1; v <= 3; v++ {
		if pos, ok := lookup[v]; ok {
			//fmt.Printf("v: %d | pos: %d | ways[pos]: %d\n", v, pos, ways[pos])
			res += int64(ways[pos])
		}
	}
	return res
}
