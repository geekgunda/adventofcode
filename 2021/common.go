package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

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

func readFileAsNumbersCsv() ([]int, error) {
	lines, err := readFileAsStrings()
	if err != nil {
		return nil, err
	}
	var res []int
	for _, l := range lines {
		numArr := strings.Split(l, ",")
		for _, num := range numArr {
			n, e := strconv.Atoi(num)
			if e != nil {
				fmt.Errorf("Error parsing num: %v", e)
			}
			res = append(res, n)
		}
	}
	return res, nil
}

func logResult(day, part int, msg string, ans interface{}) {
	fmt.Printf("Day %d | Part %d | %s: %v\n", day, part, msg, ans)
}
