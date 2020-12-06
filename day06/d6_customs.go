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
	count := GetCustomsAnswerCount(lines)
	log.Printf("{Part 1): Total Customs Answer count: %d", count)
}

func GetCustomsAnswerCount(input []string) int {
	var count int
	groupAns := make(map[rune]bool)
	for _, ans := range input {
		if len(ans) == 0 {
			count += len(groupAns)
			//log.Printf("Group Answers: %#v | Count: %d", groupAns, count)
			groupAns = map[rune]bool{}
			continue
		}
		for _, r := range ans {
			groupAns[r] = true
		}
	}
	if len(groupAns) > 0 {
		count += len(groupAns)
	}
	return count
}
