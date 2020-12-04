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
	re = regexp.MustCompile(`(\w+):(#*\w+)\s*`)
}

func main() {
	inputFile := "../input/d4.txt"
	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	lines = lines[:len(lines)-1]
	count := GetValidPassportCount(lines, false)
	log.Printf("(Part 1) Valid passport count: %d", count)
	count = GetValidPassportCount(lines, true)
	log.Printf("(Part 2) Valid passport count: %d", count)
}

func GetValidPassportCount(lines []string, strictCheck bool) int {
	var passportFields string
	var count int
	for _, l := range lines {
		if len(l) == 0 {
			//fmt.Println("New line found: ", passportFields)
			m := re.FindAllStringSubmatch(passportFields, -1)
			if CheckIfValidDoc(m, strictCheck) {
				//fmt.Println("This record is valid")
				count++
			}
			passportFields = "" // reset
		} else {
			passportFields += " " + l
		}
	}
	if len(passportFields) != 0 {
		//fmt.Println("New line found: ", passportFields)
		m := re.FindAllStringSubmatch(passportFields, -1)
		if CheckIfValidDoc(m, strictCheck) {
			//fmt.Println("This record is valid")
			count++
		}
	}
	return count
}

func CheckIfValidDoc(matches [][]string, strictCheck bool) bool {
	var validFieldCount int
	var hasCountryCode bool
	for _, match := range matches {
		switch match[1] {
		case "byr":
			if !strictCheck {
				validFieldCount++
				continue
			}
			yr, err := strconv.Atoi(match[2])
			if err != nil {
				fmt.Println("Error parsing byr: ", err)
			} else {
				if yr >= 1920 && yr <= 2002 {
					validFieldCount++
				}
			}
		case "iyr":
			if !strictCheck {
				validFieldCount++
				continue
			}
			yr, err := strconv.Atoi(match[2])
			if err != nil {
				fmt.Println("Error parsing byr: ", err)
			} else {
				if yr >= 2010 && yr <= 2020 {
					validFieldCount++
				}
			}
		case "eyr":
			if !strictCheck {
				validFieldCount++
				continue
			}
			yr, err := strconv.Atoi(match[2])
			if err != nil {
				fmt.Println("Error parsing byr: ", err)
			} else {
				if yr >= 2020 && yr <= 2030 {
					validFieldCount++
				}
			}
		case "hgt":
			if !strictCheck {
				validFieldCount++
				continue
			}
			if strings.HasSuffix(match[2], "cm") {
				cm, err := strconv.Atoi(strings.TrimSuffix(match[2], "cm"))
				if err != nil {
					fmt.Println("Error parsing byr: ", err)
				} else {
					if cm >= 150 && cm <= 193 {
						validFieldCount++
					}
				}
			}
			if strings.HasSuffix(match[2], "in") {
				in, err := strconv.Atoi(strings.TrimSuffix(match[2], "in"))
				if err != nil {
					fmt.Println("Error parsing byr: ", err)
				} else {
					if in >= 59 && in <= 76 {
						validFieldCount++
					}
				}
			}
		case "hcl":
			if !strictCheck {
				validFieldCount++
				continue
			}
			if strings.HasPrefix(match[2], "#") {
				if len(match[2]) == 7 { // including "#"
					trimmed := strings.Trim(match[2], "#0123456789abcdef")
					if len(trimmed) == 0 {
						validFieldCount++
					}
				}
			}
		case "ecl":
			if !strictCheck {
				validFieldCount++
				continue
			}
			switch match[2] {
			case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
				validFieldCount++
			}
		case "pid":
			if !strictCheck {
				validFieldCount++
				continue
			}
			if len(match[2]) == 9 {
				trimmed := strings.Trim(match[2], "0123456789")
				if len(trimmed) == 0 {
					validFieldCount++
				}
			}
		case "cid":
			validFieldCount++
			hasCountryCode = true
		}
	}
	if validFieldCount < 7 {
		//fmt.Println("Invalid record (<7): ", matches)
		return false
	}
	if validFieldCount >= 8 {
		return true
	}
	if !hasCountryCode {
		return true
	}
	//fmt.Println("Invalid record (=7): ", matches)
	return false
}
