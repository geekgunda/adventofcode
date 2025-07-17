//go:build ignore

package main

import (
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var ruleMap map[int][]int

func main() {
	input := "5.txt"
	bytes, err := os.ReadFile(input)
	if err != nil {
		log.Fatalf("Error reading file [%v]: %v", input, err)
	}
	var rules [][]int
	var updates [][]int
	for line := range strings.Lines(string(bytes)) {
		l := strings.Trim(line, "\n")
		// Ignore the empty line
		if len(l) == 0 {
			continue
		}
		// Parse page ordering rules
		if parts := strings.Split(l, "|"); len(parts) > 1 {
			var rule []int
			for _, part := range parts {
				if p, err := strconv.Atoi(part); err != nil {
					log.Fatalf("Error parsing line: [%v] | part: [%v] | err: %v", l, part, err)
				} else {
					rule = append(rule, p)
				}
			}
			rules = append(rules, rule)
		}
		// Parse update lists
		if parts := strings.Split(l, ","); len(parts) > 1 {
			var update []int
			for _, part := range parts {
				if p, err := strconv.Atoi(part); err != nil {
					log.Fatalf("Error parsing line: [%v] | part: [%v] | err: %v", l, part, err)
				} else {
					update = append(update, p)
				}
			}
			updates = append(updates, update)
		}
	}
	ruleMap = flattenRules(rules)

	var sumVerifiedMiddlePages, sumFixedMiddlePages int
	for k, update := range updates {
		log.Printf("Checking update at [%v]: %#v", k, update)
		//Check if updates are correct as per rules
		if verifyUpdate(update, ruleMap) {
			middle := len(update) / 2
			log.Printf("Update [%v] is valid | middle elem is [%v]: %v", k, middle, update[middle])
			sumVerifiedMiddlePages += update[middle]
		} else {
			// Fix them if the order is incorrect
			slices.SortFunc(update, func(i, j int) int {
				if slices.Contains(ruleMap[i], j) {
					return -1
				}
				return 1
			})
			middle := len(update) / 2
			log.Printf("Fixed update: %#v | middle elem is [%v]: %v", update, middle, update[middle])
			sumFixedMiddlePages += update[middle]
		}
	}
	log.Printf("Part 1 | Sum of middle pages: [%v]", sumVerifiedMiddlePages)
	log.Printf("Part 2 | Sum of fixed pages: [%v]", sumFixedMiddlePages)
}

func verifyUpdate(update []int, rules map[int][]int) bool {
	for i := 1; i < len(update); i++ {
		// Check that all pages BEFORE this aren't part of the ruleMap for this page number
		for j := i - 1; j < i; j++ {
			if slices.Contains(rules[update[i]], update[j]) {
				return false
			}
		}
	}
	return true
}

// Flatten all page ordering rules into a single map
func flattenRules(rules [][]int) map[int][]int {
	// ruleMap gives the list of pages that should be printed BEFORE the page number as map key
	ruleMap := make(map[int][]int)
	for _, rule := range rules {
		r1, r2 := rule[0], rule[1]
		ruleMap[r1] = append(ruleMap[r1], r2)
	}
	return ruleMap
}
