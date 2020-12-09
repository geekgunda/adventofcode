package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	inputFile := "../input/d6.txt"
	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	lines = lines[:len(lines)-1]
	anyCount, everyCount := GetCustomsAnswerCount(lines)
	log.Printf("(Part 1): Total Customs answer count (anyone answered): %d", anyCount)
	log.Printf("(Part 2): Total Customs answer count (everyone answered): %d", everyCount)
}

func GetCustomsAnswerCount(input []string) (int, int) {
	var anyCount, everyCount, groupCount int
	groupAns := make(map[rune]int)
	for _, ans := range input {
		if len(ans) == 0 {
			anyCount += len(groupAns)
			for _, ct := range groupAns {
				if ct == groupCount {
					everyCount++
				}
			}
			//log.Printf("Group Answers: %#v | Count: %d", groupAns, count)
			groupAns = map[rune]int{}
			groupCount = 0
			continue
		}
		for _, r := range ans {
			if _, ok := groupAns[r]; ok {
				groupAns[r]++
			} else {
				groupAns[r] = 1
			}
		}
		groupCount++
	}
	if len(groupAns) > 0 {
		anyCount += len(groupAns)
		for _, ct := range groupAns {
			if ct == groupCount {
				everyCount++
			}
		}
	}
	return anyCount, everyCount
}
