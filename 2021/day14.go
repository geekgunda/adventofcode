package main

import (
	"fmt"
	"strings"
)

func day14() error {
	input, err := readFileAsStrings()
	if err != nil {
		return err
	}
	//count := measurePolymerGrowth(input, 10)
	count := growPolymer(input, 10)
	logResult(14, 1, "Count after 10 steps: ", count)
	count = growPolymer(input, 40)
	logResult(14, 2, "Count after 40 steps: ", count)
	return nil
}

// Naive approach of trying to actually maintain the string and measure growth
func measurePolymerGrowth(input []string, numSteps int) int64 {
	template, rules := parsePolymerizationInput(input)
	for i := 0; i < numSteps; i++ {
		template = polymerize(template, rules)
		fmt.Println("step: ", i+1, " length: ", len(template))
	}
	return countElementsDifference(template)
}

func countElementsDifference(template string) int64 {
	var mcElement, lcElement string
	counts := make(map[string]int64)
	for _, elem := range template {
		if _, ok := counts[string(elem)]; ok {
			counts[string(elem)]++
		} else {
			counts[string(elem)] = int64(1)
		}
		if ct, ok := counts[mcElement]; ok {
			if ct < counts[string(elem)] {
				mcElement = string(elem)
			}
		} else {
			mcElement = string(elem)
		}
		if ct, ok := counts[lcElement]; ok {
			if ct > counts[string(elem)] {
				lcElement = string(elem)
			}
		} else {
			lcElement = string(elem)
		}
	}
	fmt.Println("counts: ", counts)
	return counts[mcElement] - counts[lcElement]
}

func polymerize(template string, rules map[string]string) string {
	res := string(template[0])
	for i := 0; i < len(template)-1; i++ {
		if element, ok := rules[template[i:i+2]]; ok {
			res += element
		}
		res += string(template[i+1])
	}
	return res
}

func parsePolymerizationInput(input []string) (string, map[string]string) {
	rules := make(map[string]string)
	template := input[0]
	for _, line := range input[2:] {
		parts := strings.Split(line, " -> ")
		if len(parts) != 2 {
			fmt.Println("Invalid rule: ", line)
		}
		rules[parts[0]] = parts[1]
	}
	return template, rules
}

func addCountToMap(counts map[string]int64, key string, count int64) {
	if _, ok := counts[key]; ok {
		counts[key] += count
	} else {
		counts[key] = count
	}
}

func maxElementDiff(counts map[string]int64, template string) int64 {
	var min, max int64
	var maxEle, minEle string
	elementCount := make(map[string]int64)
	// just count second character of each pair to avoid duplicates
	for pair, ct := range counts {
		addCountToMap(elementCount, string(pair[1]), ct)
	}
	// the first element of the template needs to be counted!
	// It would never change, since all additions happen in the middle
	// This is very clever and I did not come up with this by myself!
	addCountToMap(elementCount, string(template[0]), int64(1))
	for ele, ct := range elementCount {
		if min == 0 {
			min = ct
		}
		if ct < min {
			min = ct
			minEle = ele
		}
		if ct > max {
			max = ct
			maxEle = ele
		}
	}
	fmt.Println("Max: ", maxEle, max, " | Min: ", minEle, min)
	return max - min
}

// V2 function that uses rules to update counts instead
func growPolymer(input []string, steps int) int64 {
	template, rules := parsePolymerizationInput(input)
	// split the template into pairs and count their occurence
	pairCounts := pairCountsInTemplate(template)
	// Iterate and increment counts for each pair found in the rules
	for i := 0; i < steps; i++ {
		//fmt.Println(i, " pair counts: ", pairCounts)
		pairCounts = applyPolymerGrowthRules(pairCounts, rules)
	}
	//fmt.Println(steps, " pair counts: ", pairCounts)
	// finally calculate element count from pair count and calculate max difference in counts
	return maxElementDiff(pairCounts, template)
}

func applyPolymerGrowthRules(counts map[string]int64, rules map[string]string) map[string]int64 {
	// We need to store new counts in a new map
	// This is because additions in current step shouldn't be counted for subsequent insertions within this step
	nextCounts := make(map[string]int64)
	for p, c := range rules {
		if ct, ok := counts[p]; ok {
			// create new pairs
			p1, p2 := string(p[0])+c, c+string(p[1])
			// add this count to each new pair
			addCountToMap(nextCounts, p1, ct)
			addCountToMap(nextCounts, p2, ct)
		}
	}
	return nextCounts
}

func pairCountsInTemplate(template string) map[string]int64 {
	counts := make(map[string]int64)
	for i := 0; i < len(template)-1; i++ {
		pair := string(template[i : i+2])
		addCountToMap(counts, pair, int64(1))
	}
	return counts
}
