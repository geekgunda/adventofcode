package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(`(\d+)-(\d+)\s(\w+):\s(\w+)`)
}

func main() {
	inputFile := "../input/d2.txt"
	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	lines = lines[:len(lines)-1]
	validCount, err := GetValidPasswordCount(lines, false)
	if err != nil {
		log.Fatalf("Error processing: %v", err)
	}
	log.Printf("(Part1) Valid password count: %d", validCount)
	if validCount, err = GetValidPasswordCount(lines, true); err != nil {
		log.Fatalf("Error processing part2: %v", err)
	}
	log.Printf("(Part2) Valid password count: %d", validCount)

}

func GetValidPasswordCount(lines []string, policyPart2 bool) (int, error) {
	var ct, pos1, pos2 int
	var err error
	for _, line := range lines {
		p := re.FindStringSubmatch(line)
		if len(p) != 5 {
			return 0, fmt.Errorf("Invalid pattern with p: %v", len(p))
		}
		min, max, key, password := p[1], p[2], p[3], p[4]
		if pos1, err = strconv.Atoi(min); err != nil {
			return 0, err
		}
		if pos2, err = strconv.Atoi(max); err != nil {
			return 0, err
		}
		if policyPart2 == false {
			if VerifyPolicyFromPart1(pos1, pos2, key, password) {
				ct++
			}
		} else {
			if VerifyPolicyFromPart2(pos1, pos2, key, password) {
				ct++
			}
		}
	}
	return ct, nil
}

func VerifyPolicyFromPart1(pos1, pos2 int, key, password string) bool {
	rep := strings.Count(password, key)
	if rep >= pos1 && rep <= pos2 {
		return true
	}
	return false
}

func VerifyPolicyFromPart2(pos1, pos2 int, key, password string) bool {
	ct := 0
	if pos1 <= len(password) && string(password[pos1-1]) == key {
		ct++
	}
	if pos2 <= len(password) && string(password[pos2-1]) == key {
		ct++
	}
	return (ct == 1)
}
