package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var re1, re2 *regexp.Regexp

func init() {
	re1 = regexp.MustCompile(`(\w+) (\w+) bags contain`)
	re2 = regexp.MustCompile(`(\d+) (\w+) (\w+) (bags|bag)`)
}

type BagRule struct {
	Parent string
	Color  string
	Count  int
}

func main() {
	inputFile := "../input/d7.txt"
	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	lines = lines[:len(lines)-1]
	parentCount, childCount := FindMatches(lines, "shiny gold")
	log.Printf("(Part 1) Count of bags containing at least one shiny gold bag: %d", parentCount)
	log.Printf("(Part 2) Count of bags required inside shiny gold bag: %d", childCount)
}

func FindMatches(rules []string, itemToMatch string) (int, int) {
	var parentCount, childCount int
	validParents := make(map[string]bool)
	childrenOf := make(map[string][]BagRule)
	containedBy := make(map[string][]BagRule)
	for _, rule := range rules {
		//log.Printf("Rule: %s", rule)
		parent := re1.FindStringSubmatch(rule)
		if len(parent) != 3 {
			log.Fatalf("Invalid parent match: %#v", parent)
		}
		parentColor := parent[1] + " " + parent[2]
		matches := re2.FindAllStringSubmatch(rule, -1)
		for _, match := range matches {
			//log.Printf("\tMatch: %s", match[2]+" "+match[3])
			if len(match) != 5 {
				log.Fatalf("Invalid child match: %#v", match)
			}
			childColor := match[2] + " " + match[3]
			nestedCount, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatalf("Failed parsing bag count: %v | Rule: %s", err, rule)
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
				log.Printf("Found match: %s", parentColor)
				parentCount++
				validParents[parentColor] = true
			}
		}
	}
	log.Printf("Total containedBy: %d", len(containedBy))
	var nestedLevel int
	foundNewMatch := true
	for foundNewMatch {
		foundNewMatch = false
		nestedLevel++
		log.Printf("Nested level: %d | count: %d", nestedLevel, parentCount)
		for color, _ := range validParents {
			if parents, ok := containedBy[color]; ok {
				//log.Printf("New nested parent for %s | count: %d", color, len(parents))
				for _, p := range parents {
					if _, ok := validParents[p.Parent]; !ok {
						//log.Printf("Found nested match: %s", p.Parent)
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

	return parentCount, childCount
}

func getNestedCount(childrenBags []BagRule, childrenOf map[string][]BagRule) int {
	finalCount := 0
	for _, c := range childrenBags {
		bagCount := c.Count
		log.Printf("Looking for children of: %s | bagCount: %d", c.Color, bagCount)
		nestedChildren, ok := childrenOf[c.Color]
		if !ok {
			log.Printf("No children for %s | BagCount: %d", c.Color, bagCount)
			finalCount += bagCount
			continue
		}
		finalCount += bagCount + bagCount*getNestedCount(nestedChildren, childrenOf)
		log.Printf("Final count after %s is %d", c.Color, finalCount)
	}
	return finalCount
}
