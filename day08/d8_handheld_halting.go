package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(`(\w+) (\W\w+)`)
}

func main() {
	inputFile := "../input/d8.txt"
	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	lines = lines[:len(lines)-1]
	res := FindLoop(lines)
	log.Printf("(Part 1) Accumulator value before loop: %d", res)
}

func FindLoop(commands []string) int {
	var res int
	lookup := make(map[int]bool)
	for i := 0; i < len(commands); {
		if _, ok := lookup[i]; ok {
			log.Printf("Loop detected at pos: %d", i)
			break
		}
		lookup[i] = true
		m := re.FindStringSubmatch(commands[i])
		if len(m) != 3 {
			log.Printf("Invalid pattern: %s", commands[i])
			continue
		}
		step, err := strconv.Atoi(m[2])
		if err != nil {
			log.Printf("Error parsing step: %v", err)
			continue
		}
		switch m[1] {
		case "acc":
			res += step
			i++
		case "jmp":
			i += step
		case "nop":
			i++
		}
		log.Printf("Processed command: %s | step: %v | pos: %d | accumulator: %d", m[1], step, i, res)
	}
	return res
}
