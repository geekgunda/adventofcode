package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type BagRule struct {
	Parent string
	Color  string
	Count  int
}

func day7() error {
	lines, err := readFileAsStrings()
	if err != nil {
		return err
	}
	parentCount, childCount, err := FindMatches(lines, "shiny gold")
	if err != nil {
		return err
	}
	logResult(7, 1, "Count of bags containing at least one shiny gold bag", parentCount)
	logResult(7, 2, "Count of bags required inside shiny gold bag", childCount)
	return nil
}

func FindMatches(rules []string, itemToMatch string) (parentCount, childCount int, err error) {
	re1 := regexp.MustCompile(`(\w+) (\w+) bags contain`)
	re2 := regexp.MustCompile(`(\d+) (\w+) (\w+) (bags|bag)`)
	validParents := make(map[string]bool)
	childrenOf := make(map[string][]BagRule)
	containedBy := make(map[string][]BagRule)
	for _, rule := range rules {
		parent := re1.FindStringSubmatch(rule)
		if len(parent) != 3 {
			err = fmt.Errorf("Invalid parent match: %#v", parent)
			return
		}
		parentColor := parent[1] + " " + parent[2]
		matches := re2.FindAllStringSubmatch(rule, -1)
		for _, match := range matches {
			if len(match) != 5 {
				err = fmt.Errorf("Invalid child match: %#v", match)
				return
			}
			childColor := match[2] + " " + match[3]
			nestedCount, er := strconv.Atoi(match[1])
			if err != nil {
				err = fmt.Errorf("Failed parsing bag count: %v | Rule: %s", er, rule)
				return
			}
			newBagRule := BagRule{
				Parent: parentColor,
				Color:  childColor,
				Count:  nestedCount,
			}
			if _, ok := containedBy[childColor]; ok {
				containedBy[childColor] = append(containedBy[childColor], newBagRule)
			} else {
				containedBy[childColor] = []BagRule{newBagRule}
			}
			if _, ok := childrenOf[parentColor]; ok {
				childrenOf[parentColor] = append(childrenOf[parentColor], newBagRule)
			} else {
				childrenOf[parentColor] = []BagRule{newBagRule}
			}
			if childColor == itemToMatch {
				//fmt.Printf("Found match: %s\n", parentColor)
				parentCount++
				validParents[parentColor] = true
			}
		}
	}
	//fmt.Printf("Total containedBy: %d\n", len(containedBy))
	var nestedLevel int
	foundNewMatch := true
	for foundNewMatch {
		foundNewMatch = false
		nestedLevel++
		//fmt.Printf("Nested level: %d | count: %d\n", nestedLevel, parentCount)
		for color, _ := range validParents {
			if parents, ok := containedBy[color]; ok {
				//fmt.Printf("New nested parent for %s | count: %d\n", color, len(parents))
				for _, p := range parents {
					if _, ok := validParents[p.Parent]; !ok {
						//fmt.Printf("Found nested match: %s\n", p.Parent)
						validParents[p.Parent] = true
						parentCount++
						foundNewMatch = true
					}
				}
			}
		}
	}
	childrenBags := childrenOf[itemToMatch]
	childCount = getNestedCount(childrenBags, childrenOf)
	return

}

func getNestedCount(childrenBags []BagRule, childrenOf map[string][]BagRule) int {
	finalCount := 0
	for _, c := range childrenBags {
		bagCount := c.Count
		//fmt.Printf("Looking for children of: %s | bagCount: %d\n", c.Color, bagCount)
		nestedChildren, ok := childrenOf[c.Color]
		if !ok {
			//fmt.Printf("No children for %s | BagCount: %d\n", c.Color, bagCount)
			finalCount += bagCount
			continue
		}
		finalCount += bagCount + bagCount*getNestedCount(nestedChildren, childrenOf)
		//fmt.Printf("Final count after %s is %d\n", c.Color, finalCount)
	}
	return finalCount
}
