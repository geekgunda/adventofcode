package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

var re1, re2 *regexp.Regexp

func init() {
	re1 = regexp.MustCompile(`(\w+) (\w+) bags contain`)
	re2 = regexp.MustCompile(`(\d+) (\w+) (\w+) (bags|bag)`)
}

type BagRule struct {
	parent   string
	children []string
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
	count := FindMatchingParentBags(lines, "shiny gold")
	log.Printf("(Part 1) Count of bags containing at least one shiny gold bag: %d", count)
}

func FindMatchingParentBags(rules []string, itemToMatch string) int {
	var count int
	validParents := make(map[string]bool)
	containedBy := make(map[string][]string)
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
			if _, ok := containedBy[childColor]; ok {
				containedBy[childColor] = append(containedBy[childColor], parentColor)
			} else {
				containedBy[childColor] = []string{parentColor}
			}
			if childColor == itemToMatch {
				log.Printf("Found match: %s", parentColor)
				count++
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
		log.Printf("Nested level: %d | count: %d", nestedLevel, count)
		for color, _ := range validParents {
			if parents, ok := containedBy[color]; ok {
				log.Printf("New nested parent for %s | count: %d", color, len(parents))
				for _, p := range parents {
					if _, ok := validParents[p]; !ok {
						log.Printf("Found nested match: %s", p)
						validParents[p] = true
						count++
						foundNewMatch = true
					}
				}
			}
		}
	}

	return count
}
