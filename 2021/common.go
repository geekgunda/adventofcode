package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// Read input file as a slice of string
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

// Read input file as a slice of int64
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

// Read input file as an IntSlice for easy sorting later
func readFileAsIntSlice() (sort.IntSlice, error) {
	lines, err := readFileAsStrings()
	if err != nil {
		return nil, err
	}
	input := make(sort.IntSlice, len(lines))
	for i, l := range lines {
		num, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		input[i] = num
	}
	return input, nil
}

// Read input file as a multi line comma separated numbers
func readFileAsNumbersCsv() ([]int, error) {
	lines, err := readFileAsStrings()
	if err != nil {
		return nil, err
	}
	var res []int
	for _, l := range lines {
		numbers, err := readLineAsNumbersCsv(l)
		if err != nil {
			return nil, fmt.Errorf("Error parsing num: %v", err)
		}
		res = append(res, numbers...)
	}
	return res, nil
}

func readLineAsNumbersCsv(line string) ([]int, error) {
	var res []int
	numArr := strings.Split(line, ",")
	for _, num := range numArr {
		n, e := strconv.Atoi(num)
		if e != nil {
			return nil, fmt.Errorf("Error parsing num: %v", e)
		}
		res = append(res, n)
	}
	return res, nil
}

// Read input file as a 2-D integer array
func readFileAs2DIntArray() ([][]int, error) {
	lines, err := readFileAsStrings()
	if err != nil {
		return nil, err
	}
	res := make([][]int, len(lines))
	for i, line := range lines {
		res[i] = make([]int, len(line))
		for j, r := range line {
			n, e := strconv.Atoi(string(r))
			if e != nil {
				return nil, fmt.Errorf("Error parsing num: %v", e)
			}
			res[i][j] = n
		}
	}
	return res, nil
}

// Log the result in a standard format, listing the day and part of problem
func logResult(day, part int, msg string, ans interface{}) {
	fmt.Printf("Day %d | Part %d | %s: %v\n", day, part, msg, ans)
}
