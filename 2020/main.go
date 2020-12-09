package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("input", "1.txt", "input file relative path")
var problemNum = flag.Int("problem", 1, "problem number to run")

func main() {
	flag.Parse()
	fmt.Printf("Solving problem [%d] with input file at: %s\n", *problemNum, *inputFile)
	if err := solveProblem(*problemNum); err != nil {
		fmt.Printf("Error while solving problem: %v\n", err)
	}
}

func solveProblem(num int) (err error) {
	switch num {
	case 1:
		err = day1()
	}
	return err
}

func readFileAsStrings() ([]string, error) {
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return nil, err
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	lines = lines[:len(lines)-1]
	return lines, nil
}

func readFileAsInts64() ([]int64, error) {
	lines, err := readFileAsStrings()
	if err != nil {
		return nil, err
	}
	input := make([]int64, len(lines))
	for i, l := range lines {
		num, err := strconv.ParseInt(l, 10, 64)
		if err != nil {
			return nil, err
		}
		input[i] = num
	}
	return input, nil
}

func day1() error {
	numbers, err := readFileAsInts64()
	if err != nil {
		return err
	}
	result := int64(2020)
	product, err := FindMatchingPair(numbers, result)
	if err != nil {
		return err
	}
	fmt.Printf("Day 1 | Part 1 | Product (2) is: %d\n", product)
	if product, err = FindMatchingTriplet(numbers, result); err != nil {
		return err
	}
	fmt.Printf("Day 1 | Part 2 | Product (3) is: %d\n", product)
	return nil
}
