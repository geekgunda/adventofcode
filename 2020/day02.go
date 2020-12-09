package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func day2() error {
	input, err := readFileAsStrings()
	if err != nil {
		return err
	}
	validCount, err := GetValidPasswordCount(input, false)
	if err != nil {
		return err
	}
	logResult(2, 1, "Valid password count", validCount)
	if validCount, err = GetValidPasswordCount(input, true); err != nil {
		return err
	}
	logResult(2, 2, "Valid password count", validCount)
	return nil
}

func GetValidPasswordCount(lines []string, policyPart2 bool) (int, error) {
	var ct, pos1, pos2 int
	var err error
	re := regexp.MustCompile(`(\d+)-(\d+)\s(\w+):\s(\w+)`)
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
